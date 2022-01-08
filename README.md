# API Test
______

## Overview

The financial world comes with a lot of crazy file formats which don't match up with today's standards of machine-readable formats such as JSON, Protobuf, or even XML.
The goal here is simple: __write a parser for a colon-delimited file format__. 
While there are actual file formats for payments used within the FinTech industry, we’re going to use a simplified version to reduce complexity.
We have provided test data which should be used to test your logic.

The colon-delimited format works as follows:

```
{SENDER NAME:SENDER ACCOUNT NUMBER: SENDER ROUTING NUMBER: SENDER NOTE: RECIPIENT ACCOUNT NUMBER: RECIPIENT ROUTING NUMBER}
```

- SENDER NAME: ASCII string, max 100 characters, cannot contain colon, required
- SENDER ACCOUNT NUMBER: positive integer, maximum 20 characters, required
- SENDER ROUTING NUMBER: positive integer, maximum 20 characters, required
- SENDER NOTE: ASCII string, cannot contain colon, maximum 100 characters
- RECIPIENT ACCOUNT NUMBER: positive integer, maximum 20 characters, required
- RECIPIENT ROUTING NUMBER: positive integer, maximum 20 characters, required

The combination of sender's account number, and sender's routing number must not be the same as the recipient account number and recipient routing number.

For the purposes of this example, there will be one payment per file. This payment may be spread across multiple lines.

All payment information is contained within the curly braces { }

## Folder Structure Notes

`testdata` contains a number of files that can be used for testing against. Do not change these, but feel free to add more.

`helpers_test.go` contains helper functions that can be edited to suit your testing style. These will load all files in testdata into a test table.

`parse_test.go` contains a sample stub for initiating tests for your logic

`parse.go` should contain the logic for parsing the colon-delimited file format

Feel free to add more files, and tests where needed. As mentioned above, we would ask that you do not edit the sample test cases unless necessary.

## Acceptance Criteria

Your project will be judged based on the following:

- Do all the test cases pass?
- Have you considered edge cases that may not be covered by the tests?
- Does to code adhere to idiomatic Go practices?
- Time Complexity
- Space Complexity

## Running the tests

Whilst in the `api-home` folder, run the following command
`go test -v ./...`

## Summary

Reserve Trust engineers will review your code, and use it as a basis for discussion in your next interview stages.

We ask that you limit your time to an absolute maximum of 2 hours. If the exercise is not complete, do not worry!

We value your time, and would rather discuss what is missing, than have you spend all day working on this exercise.