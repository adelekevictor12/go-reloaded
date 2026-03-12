# Go-Reloaded

A text modification tool written in Go that reads an input file, applies a set of transformation rules, and writes the result to an output file.

## Usage
```bash
go run . <input_file> <output_file>
```

The program takes exactly two arguments:

- `input_file` – the file containing the text to be modified.
- `output_file` – the file where the modified text will be written.

## Supported Modifications

### Hexadecimal and Binary Conversion

- `(hex)` – Replaces the preceding word (a hexadecimal number) with its decimal equivalent.
  - `"1E (hex) files were added"` → `"30 files were added"`
- `(bin)` – Replaces the preceding word (a binary number) with its decimal equivalent.
  - `"It has been 10 (bin) years"` → `"It has been 2 years"`

### Case Conversion

- `(up)` – Converts the preceding word to uppercase.
  - `"Ready, set, go (up) !"` → `"Ready, set, GO!"`
- `(low)` – Converts the preceding word to lowercase.
  - `"I should stop SHOUTING (low)"` → `"I should stop shouting"`
- `(cap)` – Capitalizes the preceding word (title case).
  - `"Welcome to the Brooklyn bridge (cap)"` → `"Welcome to the Brooklyn Bridge"`

Each of these case tags also supports an optional number argument to apply the conversion to multiple preceding words:

- `(up, <number>)` / `(low, <number>)` / `(cap, <number>)`
  - `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`

### Punctuation Formatting

Punctuation marks (`. , ! ? : ;`) are moved to be directly after the preceding word with a space before the next word:

- `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`

Groups of punctuation like `...` or `!?` are kept together and attached to the preceding word:

- `"I was thinking ... You were right"` → `"I was thinking... You were right"`

### Quotation Marks

Single quotes (`'`) are paired and placed directly around the enclosed word(s) with no extra spaces:

- `"I am exactly how they describe me: ' awesome '"` → `"I am exactly how they describe me: 'awesome'"`
- `"As Elton John said: ' I am the most well-known homosexual in the world '"` → `"As Elton John said: 'I am the most well-known homosexual in the world'"`

### Indefinite Article Correction

`a` is automatically changed to `an` when the next word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`:

- `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`

## Implementation Details

The program processes transformations in the following order:

1. **Tag processing** (`tags`) – Handles `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` and their numbered variants.
2. **Quote formatting** (`quote`) – Pairs single quotes and attaches them to enclosed words.
3. **Punctuation formatting** (`punc`) – Adjusts spacing around punctuation marks.
4. **Article correction** (`article`) – Converts `a` to `an` before vowel/h-initial words.

## Dependencies

- [golang.org/x/text](https://pkg.go.dev/golang.org/x/text) – Used for title-case (capitalization) via `cases.Title`.



## Author

Victor