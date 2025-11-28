# Beeflang Design Specification

**A meme language honoring the Church of Beef**

## Core Philosophy
Beef-themed keywords with religious overtones. Syntax should be humorous but make intuitive sense where possible.

---

## Agreed Keywords

- `praise` - function declaration
- `beef` - block terminator (end of function/loop/conditional)
- `feast while` - while loop
- `if` / `else` - conditionals
- `prep` - variable declaration (mutable)
- `serve` - return statement
- `genesis` - entry point (main function)
- `wrangle` - import/load a module
- `herd` - module/namespace

**Current Syntax:**
```
wrangle io

praise genesis():
   prep x = 42
   io.preach(x)

   if x > 0:
      prep y = 5
   beef

   feast while x > 0:
      io.preach(x)
      x = x - 1
   beef
beef
```

---

## Operators (Standard)

### Arithmetic
- `+` - addition
- `-` - subtraction
- `*` - multiplication
- `/` - division
- `%` - modulo

### Comparison
- `==` - equal
- `!=` - not equal
- `<` - less than
- `>` - greater than
- `<=` - less than or equal
- `>=` - greater than or equal

### Logical
- `&&` or `and` - logical and
- `||` or `or` - logical or
- `!` or `not` - logical not

### Assignment
- `=` - assignment

---

## Data Types

- **int** - integers
- **bool** - true/false
- **string** - text literals (double-quoted: `"Hello, Beef!"`)

**Type System:** Dynamically typed (inferred from value)

---

## ToDo: Turing Completeness Requirements

### 1. ~~Conditionals (if/else)~~ ✓
- [x] Keywords: `if` / `else`

### 2. ~~Variables~~ ✓
- [x] Variable declaration: `prep`
- [x] Scope rules: Block-level scoping

### 3. ~~Return Statement~~ ✓
- [x] Keyword: `serve`

### 4. ~~Operators~~ ✓
- [x] Standard operators

### 5. ~~Data Types~~ ✓
- [x] int, bool, string (dynamically typed)

### 6. ~~Entry Point~~ ✓
- [x] `genesis` - main function

### 7. Module System
- [ ] Keywords: `wrangle` (import), `herd` (module)
- [ ] Dot notation for member access: `module.function()`
- [ ] Standard library modules

### 8. Basic I/O (via stdlib)
- [ ] `io.preach()` - print/output
- [ ] `io.input()` - read input

---

## Language Details

### Scope Rules
- **Block-level scoping**: Variables are scoped to the block they're declared in
- Variables declared in a function/loop/conditional only exist within that block and nested blocks

### Statement Terminators
- **Newline-terminated**: Statements end at newlines
- No semicolons required

### Comments
- **Single-line**: `#` (Python-style)
  ```
  # This is a comment
  prep x = 5  # inline comment
  ```

### Module System
- **Import syntax**: `wrangle <module_name>`
- **Member access**: `module.member` (dot notation)
- **Module keyword**: `herd` (for defining modules)
- Modules are namespaces containing functions and values

---

## Complete Example

```beeflang
wrangle io

# Calculate factorial recursively
praise factorial(n):
   if n <= 1:
      serve 1
   beef

   prep result = n * factorial(n - 1)
   serve result
beef

praise genesis():
   io.preach("Welcome to the Church of Beef!")

   prep num = 5
   prep answer = factorial(num)
   io.preach("Factorial of 5 is:")
   io.preach(answer)  # outputs: 120
beef
```

### 4. Operators
- [ ] Arithmetic: +, -, *, /, %
- [ ] Comparison: ==, !=, <, >, <=, >=
- [ ] Logical: and, or, not
- [ ] Assignment: =

### 5. Data Types
- [ ] What types to support? (int, float, string, bool at minimum?)
- [ ] Type system approach (static/dynamic/inferred?)

### 6. Entry Point
- [ ] Main function keyword/convention
- [ ] Program execution model

### 7. Basic I/O (Optional but useful)
- [ ] Print/output function
- [ ] Input function

---

## Brainstorm 
### Potential Keywords

#### Control Flow
- `sermon` - if statement
- `heresy` - else
- `sacrifice` - break
- `repent` - continue

#### Data/Variables
- `offering` - variable declaration
- `sacred` - constant declaration
- `congregation` - object/struct
- `scripture` - string type

#### Operators
- `smite` - not/negation
- `unite` - and
- `schism` - or

#### Functions
- `blessing` - return statement

#### Module System (AGREED - moved to main spec)
- `wrangle` - import/include ✓
- `herd` - module/namespace ✓
- `banish` - delete/remove

#### Other
- `genesis` - main/entry point
- `amen` - statement terminator
- `emptiness` or `famine` - null/void value

### Other Ideas

#### Immutable vs Mutable variables
- Mutable: `prep` (v1 - CURRENT)
- Immutable: `pack` (v2 - later)