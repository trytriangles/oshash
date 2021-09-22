# oshash (Open Subtitles Hash)

## Description

This package provides both a command-line utility and a Go library for computing Open Subtitles-format hashes for files and arbitrary data.

The Open Subtitles hash is an 8-byte identifying value produced from the input data's size and from 64 kibibytes (65536 bytes) at the beginning and end of the data, and made by simply interpreting chunks of it (which may overlap) as integers and summing them. It goes without saying that this is not a substitute for a SHA256 or Keccak hash of the full data. Instead it's a very compact and very cheap-to-produce identifier useful for fast deduplication and match-narrowing. 

## CLI usage

### Options

- `-x`: show hexadecimal values (default).
- `-b`: show binary values.
- `-d`: show decimal values (64-bit integers).
- `-f`: show filename with output.
- `-pipe`: read lines from stdin and write results to stdout.
- `-sep`: specify string to use to separate fields in tabular output (default: tab).

### Example

```
> oshash parrots.jpg
f44f78d5e4a8fbee

> oshash -d parrots.jpg
17604422328474205166

> cat filesnames.txt | oshash -pipe
f44f78d5e4a8fbee
bf01ea2bi00a21f7
a189811aabf1ce20

> cat filesnames.txt | oshash -pipe -f -x -d
images/parrot.jpg   f44f78d5e4a8fbee    17604422328474205166
images/toucan.jpg   bf01ea2bi00a21f7    10242892102200212726
images/lorikeet.jpg a189811aabf1ce20    18827272012277274944

> oshash -f -sep "," "parrot.jpg" "toucan.jpg"
parrot.jpg,f44f78d5e4a8fbee
toucan.jpg,bf01ea2bi00a21f7
```

## Library usage

- `FromFile(*os.File) (uint64, err)`
- `FromFilepath(string) (uint64, err)`
- `FromBytes([]byte) (uint64, err)`

To convert the uint64 values to a hexdecimal string, use `strconv.FormatUint(oshash_value, 16)`.

## Credits

Test-suite parrot image by [edmondlafoto](https://pixabay.com/users/edmondlafoto-7913128/).

## License (MIT)

©️ 2020 [Ryan Plant](mailto:ryan@ryanplant.net)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

The software is provided "as is", without warranty of any kind, express or
implied, including but not limited to the warranties of merchantability,
fitness for a particular purpose and noninfringement. In no event shall the
authors or copyright holders be liable for any claim, damages or other
liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the
software.