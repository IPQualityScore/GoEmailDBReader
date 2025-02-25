package emaillookup

import (
	"crypto/sha256"
	"fmt"
	Reader "github.com/IPQualityScore/GoEmailDBReader/pkg/reader"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
)

type EmailLookup struct {
	Path string
}
type resultobj struct {
	CreationTime int64
	Filename     string
	Data         *Reader.Data
}

// helper function to hash a string to a big int
func hashToInt(input string) *big.Int {
	hashInt := new(big.Int)
	h := sha256.New()
	h.Write([]byte(input))
	hashInt.SetBytes(h.Sum(nil))
	return hashInt
}

// input should be a domain or email, will recursively search for the input in all tree files and get the most up to date version
func (le *EmailLookup) LookupEmail(input string) *Reader.Data {
	entries, err := os.ReadDir(le.Path)
	if err != nil {
		log.Fatal(err)
	}
	resultsEmail := []resultobj{}
	resultlockEmail := sync.Mutex{}
	resultsDomain := []resultobj{}
	resultlockDomain := sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, e := range entries {
		wg.Add(1)
		go func(db os.DirEntry) {
			defer wg.Done()
			file, err := os.Open(le.Path + db.Name())
			if err != nil {
				panic(err.Error())
			}
			readerinstance := Reader.Reader{F: file, Header: Reader.Header{}}
			err = readerinstance.Header.Deserialize(file)
			if err != nil {
				fmt.Println("got error", err)
				file.Close()
				return
			}
			if readerinstance.Header.Type == 0x01 {
				split := strings.Split(input, "@")
				if len(split) > 1 {
					input = split[len(split)-1]
				}
				data := readerinstance.ContainsOnOffset(hashToInt(input), readerinstance.Header.GetSize())
				if data != nil {
					resultlockDomain.Lock()
					resultsDomain = append(resultsDomain, resultobj{
						CreationTime: int64(readerinstance.Header.CreationTime),
						Filename:     db.Name(),
						Data:         data,
					})
					resultlockDomain.Unlock()

				}
			} else {
				data := readerinstance.ContainsOnOffset(hashToInt(input), readerinstance.Header.GetSize())
				if data != nil {
					resultlockEmail.Lock()
					resultsEmail = append(resultsEmail, resultobj{
						CreationTime: int64(readerinstance.Header.CreationTime),
						Filename:     db.Name(),
						Data:         data,
					})
					resultlockEmail.Unlock()

				}
			}

			readerinstance.Close()
		}(e)
	}
	wg.Wait()
	newestResEmail := int64(0)
	finalResultEmail := resultobj{}
	newestResDomain := int64(0)
	finalResultDomain := resultobj{}

	for _, v := range resultsEmail {
		if v.CreationTime > newestResEmail {
			newestResEmail = v.CreationTime
			finalResultEmail = v
		}
	}
	for _, v := range resultsDomain {
		if v.CreationTime > newestResDomain {
			newestResDomain = v.CreationTime
			finalResultDomain = v
		}
	}

	if newestResEmail != 0 {
		if newestResDomain != 0 {
			finalResultEmail.Data.Data = append(finalResultEmail.Data.Data, finalResultDomain.Data.Data...)
		}
		return finalResultEmail.Data
	}
	if newestResDomain != 0 {
		return finalResultDomain.Data
	}
	return nil
}
