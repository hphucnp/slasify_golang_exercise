/*
*********** DO NOT CHANGE OR COMMIT THIS FILE ***********

Special cases that you'll need to handle.

There are:
- 62 lines in total
- 6 lines of inline comments
- 34 lines of block comments
*/

main()
{
    printf(
        "// This is not a comment"
        "// This is not a comment\
        // This is not a comment" // Comment: inline
        "// This is not a comment \" not a comment as well"
        "/* This is not a comment */"
        "/* This is not a comment \
        This is not a comment */"
        "This is not a comment */"
        "// This line has a trailing comment" // Comment: \
        inline (the line-break makes this one counted as 2 lines)
    );                                        /* Comment: block
    Comment: block
    Comment: block
    */ int x = 0;                                // Comment: block, inline
    /*
    Comment: block
    */ int y = 1;                                /*
    Comment: block */

    /* NOTE:
    The next block follows the raw string literals syntax, which is a special case
    for string parsing.
    In short, the format is R"[delimiter](...)[delimiter]", where the content
    between the quotes (") are not comments.
    The delimiters before and after the parentheses must match.

    You can find more details here:
    https://en.cppreference.com/w/cpp/language/string_literal
    */

    const char * vogon_poem = R"V0G0N(
    // NOT A COMMENT:             O freddled gruntbuggly thy micturations are to me
    // NOT A COMMENT:                As plured gabbleblochits on a lurgid bee.
    // NOT A COMMENT:             Groop, I implore thee my foonting turlingdromes.
    /* NOT A COMMENT:          And hooptiously drangle me with crinkly bindlewurdles,
     * NOT A COMMENT: Or I will rend thee in the gobberwarts with my blurlecruncheon, see if I don't.
     */<<NOT A COMMENT
                    (by Prostetnic Vogon Jeltz; see p. 56/57)
    )V0G0N";
}

/* Comment: block
Comment: block \
/* Comment: block // This is NOT an inline comment
/* Comment: block
/* Comment: block
Comment: block */ // Comment: inline // Not another inline comment
//* Comment: inline
