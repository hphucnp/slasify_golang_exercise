Go Developer Evaluation Program
---

## Code Interview Process

Please follow the instructions stated below and fulfill the [objective](#Objective)

You have **seven days** to finish this program.

- 🎬 Clone the repo and create a working branch based on `main` instead of forking.
- ✍🏼 Take the `// TODO` in *main.go* as your starting point.
- 📚 Implement the core logic of the program **from scratch and by your own understanding**. You can use open source libraries **ONLY** for the less essential parts, such as logging, printing, etc.
- 🏁 [Create a PR](../../pulls) when you finish. Please attach the links to the references if your implementation is based on the ideas from other sources.
- 🇺🇳 Please communicate in English. **DO NOT** use Chinese in the workflow, especially in the code and commit messages.

If you have any questions, please feel free to [raise an issue](../../issues) in the repo. We're glad to help.

Please let me know when your PR's ready to review.

## Notice

This is not a real project. Your work won't be used by the company by any means. However, when you're working on this
program, consider it as an assignment in the real life and show your best programming practices.

Your code will be reviewed and scored by the other developers of the team where you will join.

Your work will earn higher score if:

- The requirements are implemented precisely.
- You split the objective into smaller tasks and commit them with proper commit messages.
- Your code is well-organized, well-formatted, and pleasing to read.
- The names of variables and functions are accurate and easy to understand.
- The most important functionalities are covered with meaningful and effective test cases.
- There are valuable comments in your code and PR, but only when necessary.
- Performance and multi-threading are taken into account.
- Your code is extensible so that it’s easy to support comment-counting for other languages.

---

## Objective

Implement a command-line tool in Golang that counts the number of comment lines in C/C++ source code.

All input source code files are assumably compilable. In other words, you don't need to check if it's a valid C/C++
code.

The reviewer should be able to run the tool as:

```shell
go run . testing/cpp
```

The tool must meet the following requirements:

1. Take a directory as the input argument. You can leverage the files provided in [testing/cpp](./testing/cpp) for
   verifying your implementation.
2. During execution, it will walk through all the C/C++ source code files
   (only `*.c`, `*.cpp`, `*.h`, `*.hpp`), recursively, in the given directory and count the comment lines in each file.
3. Count the inline comments (`//`) **AND** block comments (`/* … */`),
   **respectively**.
4. Cache the statistics and output them to the console. The result should be **the same** as below:

    ```text
   testing/cpp/lib_json/json_reader.cpp      total: 1992    inline: 134    block:   0
   testing/cpp/lib_json/json_tool.h          total:  138    inline:  13    block:  19
   testing/cpp/lib_json/json_value.cpp       total: 1634    inline: 111    block:  18
   testing/cpp/lib_json/json_writer.cpp      total: 1259    inline:  89    block:   0
   testing/cpp/special_cases.cpp             total:   62    inline:   6    block:  34
   testing/cpp/test_lib_json/fuzz.cpp        total:   54    inline:   5    block:   0
   testing/cpp/test_lib_json/fuzz.h          total:   14    inline:   5    block:   0
   testing/cpp/test_lib_json/jsontest.cpp    total:  430    inline:  54    block:   1
   testing/cpp/test_lib_json/jsontest.h      total:  288    inline:  52    block:   8
   testing/cpp/test_lib_json/main.cpp        total: 3971    inline: 182    block:   0
    ```

   Where:
    - `filename`: relative path (e.g. `testing/cpp/special_cases.cpp`) to the file. The filenames should be sorted by
      **ascending alphabetical order**.
    - `total`: total lines of the file, excluding the trailing line of EOF
    - `inline`: total lines of the inline comments in the file
    - `block`: total lines of the block comments in the file

6. :bulb:**Hint:** There are some special cases you'll need to handle. Please refer to the
   file [testing/cpp/special_cases.cpp](./testing/cpp/special_cases.cpp).

---

# Testing

```shell
go test ./...
```

Take [example_test.go](./example_test.go) as the example of writing your own test cases.

**DO NOT** put your Go test files into the [testing](./testing) directory, for it should contain only the source code
files that are to be counted upon.
