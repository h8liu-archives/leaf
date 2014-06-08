- The parser will parse the program into an abstract syntax tree
- The ast should have also recorded the position of each node in the source code
- Based on the syntax tree, we first need to compile all the constants
    - For compiling the constants in a package, we need to build the dependency tree of the constants first.
    - And then we can evaluate the constants based on the expressions.
- Next, we need to compile the structure types.
    - For compiling the types in a package, we also need to build the dependency tree of the types
    - This means that for each type, we will be able to calculate its size.
- After this, we know all the sizes of the types, we can allocate the memory positions of all the variables
  based on their declaration.
- Afterwards, we can now generate the body IR of each function. This is the first time that the control flow part of
  the IR comes in.

Or alternatively, we can build the function body first, by:
- without defining any constant (the query on any constant value on any package will be not-found)
- without defining any type other than built-in types (only support built-in types)

Stage 1: no named constants, no global variables, no types other than built-in types, only basic functions.
- will need some lexing and parsing
- will need some IR
- will need function calls
- will need expression parsing and code generation
- single module

Stage 2: module import

Stage 3: named constants

Stage 4: struct types

Stage 5: global variables

Stage 6: interface

Stage 7: optimizations...
