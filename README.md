# Second chapter of course

Building a fake hotel bookings web app. 

# Work in Progress in studying Golang properly

Chang Tan<br>
changtan@listerunlimited.com<br>
Two-time National Cyber League Winner (Individuals & Team Fall 2021 and Spring 2022)

<a href="https://www.udemy.com/course/building-modern-web-applications-with-go">Udemy Course which this code is derived from</a>

I am just doing this to study more advanced methods of Golang programming, specifically web servers and web applications since alot of features of Golang is not currently present despite being well documented in... uh... shittier languages like Java.

Of note is the changes in how the Goland IDE by JetBrains has chosen to intepret template files as they now require the .html to be appended to template .tmpl files to support Goland's IDE features.

I'm planning to use this as a template codebase to build custom libraries and Golang packages for...

1. Universal reverse proxies to protect web apps concealed behind a virtual private cloud with a built-in DNS resolver, input validation, support for common API methods such as REST, SOAP, JSON-RPC, and GraphQL
2. Cross-compilable mobile apps
3. Agricultural robotics libraries and semi-conscious "artificial intelligence"
4. Counteracting reverse-engineering attempts by competitors for production-level deployments

# Since this is basically the answers to the online course

I expect you to use this as a reference and not some sort of exam cheatsheet for the course that I linked above.

Furthermore, I am not going to be engaging anyone in some sort of online verbal feud over the superiority or inferiority and drawbacks of Golang (there are many but overall Golang is solid for Linux, and Unix host systems).

I made up my mind to become a Gopher three years ago and there are Googlable case studies on how Golang can be productive (or destructive, like Blackrota). But for the first time of my life, I decided to study Golang properly, to not be a asshole to someone.

# Cross compilation instructions (Linux)

GOOS=linux GOARCH=amd64 go build -o webapp.elf cmd/web/*.go

# Cross compilation instructions (Windows)
GOOS=windows GOARCH=amd64 go build -o webapp.exe cmd/web/*.go

# Cross compilation instructions (MacOS Desktop/MacBooks)
GOOS=darwin GOARCH=amd64 go build -o webapp.macho cmd/web/*.go <br>
**Alternatively, for M1 enabled MacBooks, they are finally supported in Golang 1.18, simply point to the ARM64 architecture**<br>
GOOS=darwin GOARCH=arm64 go build -o webapp.macho cmd/web/*.go
