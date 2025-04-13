# RandomNameGenerator
Simplistic program to pick a random name from a list defined in a JSON file.

**Building**
Build the program by running the _BUILD_ batch file from the command line.

**Execution**
Run the program by running _RANDOM <filename>_ where _<filename>_ specifies the path of the input data file to select from.

<filename> is expected to be in JSON format providing a list of strings as an array under a "names" tag. An example format is show below

**Example JSON File**
{
    "names": ["Alice", "Bob", "Charlie", "Diana", "Eve"]
}
