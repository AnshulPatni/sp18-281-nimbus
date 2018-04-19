# Golang Overview

Go, also commonly referred to as golang, is a programming language developed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a statically typed language with syntax loosely derived from that of C, adding garbage collection, type safety, some dynamic-typing capabilities, additional built-in types such as variable-length arrays & key-value maps, and a large standard library. The language was announced in November 2009 and is now used in some of Google’s production systems. Go’s “gc” compiler targets the Linux, OS X, FreeBSD, NetBSD, OpenBSD, Plan 9, DragonFly BSD, Solaris, and Windows operating systems and the i386, Amd64, ARM and IBM POWER processor architectures. A second compiler, gccgo, is a GCC frontend. Android support was added in version 1.4, which has since been ported to also run on iOS.

## Pros

1. **Simple and user-friendly**
If you have any developer experience working with other programming languages, the syntax of Go Lang will be easy and quick to grasp. Go code sometimes looks more readable and elegant even though the code itself is larger.

2. **Defined style rules**
Defined style rules allow easy maintenance of your code and other’s and it restricts any temptation to add your own style to the application.  This can be a bit annoying in the beginning, but the payoff is significant, as soon as you start working on the application with a big team.

3. **Strict compilation system**
If you set a variable and it is not in use or there are libraries that are not used in the code, you will get errors at the compilation stage. It can be annoying, but it also ensures that the code is free of clutter.

4. **Multiple return values**
Not the most useful thing at the development stage, but the availability to draw from the box sometimes is a positive factor for writing  code.

5. **Concurrency**
Parallel calculations are easily and gracefully performed in Go. You’ll have no problems and no headache.
6. Deploy and native compilation
It is a walk in the park to get a server running because Go compiles the application into the native binary, i.e. your application has no dependencies with any jwm, frameworks or libraries. There is a single file that can be deployed to any server. Version 1.5 allows compilation of a binary file to any platform, making the process even easier.

7. **Speed**
Request processing speed compared to the servers with other languages differs depending on applications,  but for Go it is a strong side.


### Go Lang has proved to be a better choice for the following tasks:

1. **Web applications and web servers.** Originally Go was created as a tool for fast and easy writing of web and mobile applications by a large number of developers and to provide an easy support environment for the code. Its own features, go routines and channels, only enhance its advantages when writing code.

2. **It proves to be a good pick as a stand-alone command-line application or script.** This language has everything going for it: a single executed file without any dependencies (if they are not needed), higher processing speed, compared to other applications, ability to work with outside C libraries and even to process system calls.

3. **A great alternative to parallel script writing in C/C++.** It is easier to write and deploy those scripts in Go.

Yes, Go is a young new language and it is not perfect–it is not a “silver bullet” that will solves every problem. But, it has strong and attractive features that can’t be ignored. There is only one thing that is for sure, you can’t answer the question “Is that language for me?” without trying it first.

## MongoDB driver for Go

mgo (pronounced as mango) is a reach MongoDB driver for Golang. Its API is very simple and follows standard Go idioms. We will see how it can help with building CRUD (create, read, update, delete) operations for microservice in a second, but first let’s get familiar with session management.

MongoDB is a very popular backend for writing microservices with Go. MongoDB driver for Go (mgo) is idiomatic and very easy to use. Don’t overlook curl if you are building, testing or documenting RESTful services.


#### References

* https://masterofcode.com/blog/an-overview-on-golang-programming-language

* http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/