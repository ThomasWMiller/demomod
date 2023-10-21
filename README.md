# demomod
Set up local modules and packages
Most Go textbooks show how to work with packages from the standard library as well as third-party packages available on public repositories. 

Curiously, they do not show how to work with local packages.

When first learning about Go, when exploring new ideas, or when developing new applications (code we do not want to share with the world in a public repository), we often prefer to work with local modules and packages. So, how do we do keep things to ourselves, working locally?

This repository shows how to set up local modules and packages. The code builds on Listing 5.35 from Bates and LaNou (2023, 159–160), which shows a function for copying a file. Note the structure and naming of the directories and files. Note the position and contents of the two **go.mod** files. Note that the **FileCopy** function begins with a capital letter so it can be exported and used in the calling program **run-filecopy.go**.

Note how we add documentation to packages. When located in the root of the **demomod** directory, type **go doc fileutil** to see the documentation for the **fileutil** package and type **go doc.fileutil.FileCopy** to see the documentation for the **FileCopy** function. 

 
##### Reference

Bates, Mark, and Cory LaNou. 2023. *Go Fundamentals: Gopher Guides.* Boston: Addison-Wesley. [ISBN-13: 978-0-13-791830-0] 
