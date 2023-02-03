# OTP-Generator
<i>This Go package is to generate One-Time Password for Email based Accoounts </i>
<h3>Usage</h3>
In your development environment,
<ul>
<li>Install OTP-Go: <code>go get github.com/seyiadel/otp-go/otp</code></li>
<li>Import and run <code>go mod tidy</code></li>
</ul>
Gorm(Go object relational mapper) is being added during installation so as to create any database of your choice to save arguments/keyword arguments.
Import Gorm<code>import ("gorm.io/gorm)</code> and gorm driver <code>import "gorm.io/driver/sqlite"</code> for your preferred database. Check <a href="https://gorm.io/docs/connecting_to_the_database.html"> Gorm.io driver</a>.
For Example using sqlite database and driver,

<code>conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		panic("failed to connect to database")
	}</code> 

Now you can call the package methods to validate and generate OTP.

<code>otp.Generate(conn,"user@email.com", 4)</code> as conn is the database, user@email.com is the user email, 4 is the desired otp length

<code>otp.ValidateOTP(conn, "user@email.com", 7645)</code>as conn is the database, user@email.com is the user email, 7645 is the token generated

<h2>Hurrayyyy!🤩🤩</h2>
