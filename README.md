<div class="documentation_overview">
	<h1 class="doc-title">IPQualityScore Email Validation &amp; Verification Golang DB Reader</h1>
	<div class="spacing-10"></div>
	<h2 class="text-bold headerHR" style="font-size: 1.5em;">Flat File Version 1.0</h2>
	<div class="spacing-10"></div>
	<p>
        Our flat file email validation database allows you to lookup important details about any email address using a straight forward library. Simply install the reader, download a database and instantly check email addresses against our large volume of data.
    </p>
    Click here to see the full <a href="https://www.ipqualityscore.com/documentation/email-validation-database/golang-email-database-reader" target="_blank">Golang IPQualityScore flat file email validation database documentation</a>. You may also find the information on our <a href="https://www.ipqualityscore.com/email-verification" target="_blank">email verification / email validation serivce</a> helpful.
    <h6 class="text-bold headerHR">Installation</h6>
	<div class="spacing-10"></div>
	<p>
        Installation can be achieved via go's built in package manager like such:
    </p>
    <div class="row">
		<div class="col-md-12 xsmcode">
			<pre class="highlight markdown"><code>
go get github.com/IPQualityScore/GoEmailDBReader
			</code></pre>
        </div>
    </div>
    <h6 class="text-bold headerHR">Usage</h6>
	<div class="spacing-10"></div>
	<p>
        Using our flat file database system to lookup an email address is simple:
    </p>
    <div class="row">
		<div class="col-md-12 lgcode">
            <pre class="highlight markdown"><code>
lookup := emaillookup.EmailLookup{Path: "./tree/"}

email := "noreply@ipqs.com"

record := lookup.LookupEmail(email)
if record != nil {
	fmt.Println("found data:", record)
 	for _, v := range record.Data {
		fmt.Println(v.ToString())
	}

	fraud := record.FraudScore()
	if fraud != nil {
        	fmt.Println("Fraud Score:", fraud.FraudScore)
	} else {
		fmt.Println("Could not find fraud score!")
	}
}
            </code></pre>
        </div>
    </div>
    <h6 class="text-bold headerHR">IPQSRecord Struct Fields</h6>
	<div class="spacing-10"></div>
	<p>
        Depending on which database file you receive some of these fields may be unavailable. If the field in question is unavailable in your database it will default to Golang's default value for that type.
    </p>
    <div class="row">
		<div class="col-md-12">
			<table class="table table-legend custom-tablelegend">
				<thead>
					<tr>
						<th style="width:10%">Field</th>
						<th style="width:10%">Type</th>
						<th>Description</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>record.Base().Valid</td>
						<td>bool</td>
						<td>Does this email address appear valid?</td>
					</tr>
                    <tr>
						<td>record.Base().Disposable</td>
						<td>bool</td>
						<td>Is this email suspected of belonging to a temporary or disposable mail service? Usually associated with fraudsters and scammers.</td>
					</tr>
                    <tr>
						<td>record.Base().SmtpScore</td>
						<td>byte</td>
						<td>
                            Validity score of email server's SMTP setup. Range: "0" to "3".<br><br>
                            0x00 = mail server exists, but is rejecting all mail<br>
                            0x01 = mail server exists, but is showing a temporary error<br>
                            0x02 = mail server exists, but accepts all email<br>
                            0x03 = mail server exists and has verified the email address<br>
                        </td>
					</tr>
                    <tr>
						<td>record.Base().Suspect</td>
						<td>bool</td>
						<td>
                            This value indicates if the mail server is currently replying with a temporary mail server error or if the email verification system is unable to verify the email address due to a broken SMTP handshake. This status will also be true for "catch all" email addresses as defined below. If this value is true, then we suspect the "valid" result may be tainted and there is not a guarantee that the email address is truly valid. This status is rarely true for popular mailbox providers and typically only returns as true for a small percentage of business mail servers.
                        </td>
					</tr>
                    <tr>
						<td>record.Base().CatchAll</td>
						<td>bool</td>
						<td>
                            Is this email likely to be a "catch all" where the mail server verifies all emails tested against it as valid? It is difficult to determine if the address is truly valid in these scenarios, since the email's server will not confirm the account's status.
                        </td>
					</tr>
                    <tr>
						<td>record.Base().Deliverability</td>
						<td>byte</td>
						<td>
                            How likely is this email to be delivered to the user and land in their mailbox. Values can be "high", "medium", "low" or "none".<br><br>
                            0x00 = none<br>
                            0x01 = low<br>
                            0x02 = medium<br>
                            0x03 = high<br>
                        </td>
					</tr>
                    <tr>
						<td>record.FraudScore().FraudScore</td>
						<td>int</td>
						<td>
                            The overall Fraud Score of the user based on the email's reputation and recent behavior across the IPQS threat network. Fraud Scores >= 75 are suspicious, but not necessarily fraudulent.
                        </td>
					</tr>
                    <tr>
						<td>record.Leaked().Leaked</td>
						<td>bool</td>
						<td>
                            Was this email address associated with a recent database leak from a third party? Leaked accounts pose a risk as they may have become compromised during a database breach.
                        </td>
					</tr>
                    <tr>
						<td>record.RecentAbuse().RecentAbuse/td>
						<td>bool</td>
						<td>
                            This value will indicate if there has been any recently verified abuse across our network for this email address. Abuse could be a confirmed chargeback, fake signup, compromised device, fake app install, or similar malicious behavior within the past few days.
                        </td>
					</tr>
                    <tr>
						<td>record.UserVelocity().UserVelocity/td>
						<td>bool</td>
						<td>
                            Frequency at which this email address makes legitimate purchases, account registrations, and engages in legitimate user behavior online. Values can be "high", "medium", "low", or "none". Values of "high" or "medium" are strong signals of healthy usage. New email addresses without a history of legitimate behavior will have a value as "none". This field is restricted to higher plan tiers.
                        </td>
					</tr>
                    <tr>
						<td>record.FirstSeen().FirstSeen</td>
						<td>time.Time</td>
						<td>
                            When this email address was first seen online.
                        </td>
					</tr>
                    <tr>
						<td>record.DomainCommon().DomainCommon</td>
						<td>bool</td>
						<td>
                            Is this email from common free email providers? ("gmail.com", "yahoo.com", "hotmail.com", etc.)
                        </td>
					</tr>
                    <tr>
						<td>record.DomainVelocity().DomainVelocity</td>
						<td>bool</td>
						<td>
                            Indicates the level of legitimate users interacting with the email address domain. Values can be "high", "medium", "low", or "none". Domains like "IBM.com", "Microsoft.com", "Gmail.com", etc. will have "high" scores as this value represents popular domains. New domains or domains that are not frequently visited by legitimate users will have a value as "none". This field is restricted to upgraded plans.
                        </td>
					</tr>
                    <tr>
						<td>record.DomainDisposable().DomainDisposable</td>
						<td>bool</td>
						<td>
                            Is this domain suspected of belonging to a temporary or disposable mail service? Usually associated with fraudsters and scammers.
                        </td>
					</tr>
                </tbody>
            </table>
        </div>
    </div>
</div>
