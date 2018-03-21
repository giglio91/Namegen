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
----------------- github.com/giglio91 -----------------
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
- length: Integer number that defines strings/names characters length (Default: 6).
- number: Integer number indicating the number of strings/names to be generated (Default: 10, 0 for all the possible combinations).
- random: Enable random strings/names generation (Default: true). Ignored if -number=0.
- vowelsPos: Comma-separated list of integer numbers that define the vowel positions in strings/names (Example: -vowelPos=2,4,6).

## Example
Generate 50 non-random (hence sequential-generated) strings/names with lenght of 5 chars. Second and fourth chars must be vowels, the others are consonants (since Alphabet Mode is false).
```
namegen.exe -length=5 -vowelsPos=2,4 -number=50 -random=false -alphabetmode=false
```
Output:
```
-------------------------------------------------------
---- STRINGS/NAMES GENERATOR BY Francesco Giglioli ----
----------------- github.com/giglio91 -----------------
-------------------------------------------------------

Random Mode:  false
Alphabet Mode: false
Strings/names length:  5
Vowels positions:  [2 4]
Number of possible combinations: 231525
Number of strings/names actually generated:  50 

babab
babac
babad
babaf
babag
babah
babaj
babak
babal
babam
baban
babap
babaq
babar
babas
babat
babav
babaw
babax
babay
babaz
babeb
babec
babed
babef
babeg
babeh
babej
babek
babel
babem
baben
babep
babeq
baber
babes
babet
babev
babew
babex
babey
babez
babib
babic
babid
babif
babig
babih
babij
babik
```
