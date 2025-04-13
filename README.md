# RandomNameGenerator
Simplistic program to pick a random name from a list defined in a JSON file.

**Building**<br>
Build the program by running the **BUILD** batch file from the command line.

**Execution**<br>
Run the program by running **RANDOM filename** <br>where **filename** specifies the path of the input data file to select from.
**filename** is expected to be in JSON format providing a list of strings as an array under a "names" tag. An example format is show below

**Example JSON File**<br>
{
    "names": ["Alice", "Bob", "Charlie", "Diana", "Eve"]
}
