# Project Title

rapairDigikam
A project written in golang (go) to facilitate correcting my corrupted Digikam photo manager.  After creating Digikam and providing a lot of insight about each individual photo, ... time went by... and I changed file names that had spaces in them to underscores, and also moved photo locations.  Several months later, when trying to use Digikam, I became aware that this name chanes and location changes had caused Digikam to not be able to associate its database with the new file names and locations!  

## Getting Started

If you have made the same mistake, then this will facilitate correcting the problem.  However, it does not directly read nor update your Digikam database (at this time); instead, you will need to use sqlite3 (or something similar) to produce a file that contains the key|filename for each nulled file.  Then run this program repairDigikam to produce an output associating the key:newfilename for use with sqlite3 (or something similar) to update your Digikam database.  MAKE SURE YOU MAKE A COPY OF YOUR DIGIKAM DATABASE BEFORE UNDERTAKING THIS PROJECT.  DO NOT PROCEED UNTIL YOU HAVE A COPY, OTHERWISE YOU MAY BE UNABLE TO EVER USE THIS DATABASE AGAIN.  Your copy is a safety precaution in case the work you are doing causes new and unforseen problems with your existing database.

### Prerequisites

You will need to have golang installed and running on your computer before you install and build this application.  No source files are provided, only the source code, so it must be built by you prior to running.

### Installing

The steps involved to use to correct Digikam database that has been corrupted because you either changed filenames that had spaces in it to underscore, or you moved file locations from where they originally existed when you added to Digikam.

Note: This process is not recommended for someone unfamiliar with programming.  It has only been tested on a computer running Ubuntu 17.10.

1.  Install and ensure golang is working.  Search golang on internet for insight if you are unfamiliar.
2.  Produce a copy of your Digikam database, calling it something like digikamdatabase.backup
2.  Using sqlite3 (or something similar) on your computer to allow you to work with your Digikam database, produce an output file containing an element for each bad entry in the Digikam database.  This file should contain the key|filename.  Note this is key "vertical bar" filename.
3.  Download and build this program repairDigikam.
4.  Using repairDigikam, either use the default filename provided in the program (do repairDigikam -h for help on default file names), or specify the filename you created above as input for this program.
5.  Run repairDigikam.
6.  Using sqlite3, (or something similar), using the output file produced by repairDigikam, update your Digikam database.
7.  Restart Digikam and see if it works correctly with your changes!
8.  If not, then revert back to your saved copy of your Digikam database.


End with an example of getting some data out of the system or using it for a little demo

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


