# Go Templater

This is an attempt at providing a general purpose library for writing
files based on templates

# Design choices

A template will require 2 types of entities to be really useful

1. A markup - containing all the non-semantics of the template and tags 
    for inserting data
    - This is by convention placed in files matching templates/_.*
    - Example: _index.html could define a beautiful homepage and some tags

2. Some data - containing a datatype that the markup file expects
    - This is by convention placed in files matching 
    templates/filename.json where filename matches a markup file like 
    templates/_(.*).json
    - Example: index.html.json could define data for our template above

Both of these types will probably be configurable at some point

## Template data

Is "reverse-enginered" from what fields are used in the template file. 
From this we attempt to find the fields in the template data file.
If not all fields can be found, the template can't be created. 

If the template files starts with a list, then all elements of the list
will be created. To make this meaningful each entry in the list
should specify a distinct file-location (see \ref{special_fields})

### Special fields
There are some "special" fields that template data can contain:
- "file-location": Specifies where to place the created file. If not specified,
the generated file will be created in "generated/"

