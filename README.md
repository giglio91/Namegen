# Namegen
Simple custom strings/names generator in Golang.

# Compiling
Simply execute:
```
go build namegen.go
```

# Usage
By default executing
```
namegen.exe
```
will return this output (example):
```
-------------------------------------------------------
---- STRINGS/NAMES GENERATOR BY Francesco Giglioli ----
-------------------------------------------------------

Random Mode:  true
Alphabet Mode: false
Strings/names length:  6
Vowels positions:  []
Number of possible combinations: 85766121
Number of strings/names actually generated:  10 

zvjpqn
ktdtnv
hlrkvr
lzzfmr
yjtplc
xystgx
pjpjmn
rzdpsz
dtkrwr
sqdwmk
```
## CMD flags
Help is provided by executing:
```
namegen.exe [--help|-h]
```
Available flags are:
- alphabetmode: Enable strings/names generation using whole alphabet (Default: false).
- length int: Integer number that define strings/names characters length (Default: 6).
- number float: Integer number indicating the number of strings/names to be generated (Default: 10, 0 for all the possible combinations).
- random: Enable random strings/names generation (Default: true). Ignored if -number=0.
- vowelsPos value: Comma-separated list of integer numbers that define the vowel positions in strings/names (Example: -vowelPos=2,4,6).
