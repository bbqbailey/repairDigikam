# Project Title

rapairDigikam
A project written in golang (go) to facilitate correcting my corrupted Digikam photo manager.  After creating Digikam and providing a lot of insight about each individual photo, ... time went by... and I changed file names that had spaces in them to underscores, and also moved photo locations.  Several months later, when trying to use Digikam, I became aware that these name changes and location changes had caused Digikam to be unable to associate its database with the new file names and locations!  

## Prerequisites
- You are using Digikam, and it is broken because, after creating your informaiton you have done one or both of the following: changed file names that had spaces in it to underscores; moved the location of the file.
- Your Digikam database has so many null-field problems that it would be nearly impossible to manually correct them.  If you only have a few (under 100), then it may be much simpler to manually repair your database.
- You are familiar with relational databases, and in particular, are familiar with Sqlite and sqlite3, or are willing to educate yourself to the extent necessary to identify and correct the problem in your database.
- You have created a copy of your Digikam database.  Incorrectly modifying your currently-broken database, may result in it being in a more-corrupted state than you currently have!
- You have studied the structure of your Digikam database and understand what is causing your problem, and this proglem is caused by nulls in the file location field.
- Your actions have the potential of corrupting your Digikam database and your ability to utulize it.

## Getting Started
MAKE SURE YOU MAKE A COPY OF YOUR DIGIKAM DATABASE BEFORE UNDERTAKING THIS PROJECT.  DO NOT PROCEED UNTIL YOU HAVE A COPY, OTHERWISE YOU MAY BE UNABLE TO EVER USE YOUR DATABASE AGAIN.  

If you have made the same mistake, then this will facilitate correcting the problem.  However, it does not directly read nor update your Digikam database (at this time); instead, you will need to use sqlite3 (or something similar) to produce a file that contains the key|filename for each nulled file.  Then run this program repairDigikam to produce an output associating the key:newfilename for use with sqlite3 (or something similar) to update your Digikam database.  MAKE SURE YOU MAKE A COPY OF YOUR DIGIKAM DATABASE BEFORE UNDERTAKING THIS PROJECT.  DO NOT PROCEED UNTIL YOU HAVE A COPY, OTHERWISE YOU MAY BE UNABLE TO EVER USE THIS DATABASE AGAIN.  Your copy is a safety precaution in case the work you are doing causes new and unforseen problems with your existing database.

### Prerequisites

You will need to have golang installed and running on your computer before you install and build this application.  No executable files are provided, only the source code, so the executable must be built by you prior to running.  If you are unfamiliar with building golang applications, then consult google search for golang to gain insight necessary to support this required step.  It is not required that you be a golang programmer, just that you are able to build and run the application.

### Installing

The steps involved to use to correct Digikam database that has been corrupted because you either changed filenames that had spaces in it to underscore, or you moved file locations from where they originally existed when you added to Digikam.

Note: This process is not recommended for someone unfamiliar with relational databases.  

Note: This application has only been tested on a computer running Ubuntu 17.10.

Note: This application does not modify your database; you are responsible for ensuring the output from this program is correct before you use it as an assistent to facilitate your action of repairing your database.

1.  Install and ensure golang is working.  Search golang on internet for insight if you are unfamiliar.
2.  Produce a copy of your Digikam database.
2.  Using sqlite3 (or something similar) on your computer to allow you to work with your Digikam database, produce an output file containing an element for each bad entry in the Digikam database.  This file should contain the key|filename.  Note this is key "vertical bar" filename.
3.  Download and build this program (repairDigikam) to facilitate your repair of Digikam.
4.  Using repairDigikam, either use the default filename provided in the program (do repairDigikam -h for help on default file names), or specify the filename you created above as input for this program.
5.  Run repairDigikam.
6.  Using sqlite3, (or something similar), using the output file produced by repairDigikam, update your Digikam database.
7.  Restart Digikam and see if it works correctly with your changes!
8.  If not, then revert back to your saved copy of your Digikam database.

## Running the tests

There are currently no automated tests.

## Built With

This is built using golang, with an additional plugin from: github.com/emirpasic/gods/tree/avltree for balanced binary search.

## Authors

Initial author: BBQBailey

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to github.com/emirpasic for his avltree balanced binary tree work!


