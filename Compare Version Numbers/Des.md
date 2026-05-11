# Compare Version Numbers

**Difficulty:** Medium

## Problem Statement

Given two version strings, `version1` and `version2`, compare them.

A version string consists of revisions separated by dots (`.`). The value of each revision is its integer conversion, ignoring leading zeros.

To compare version strings, compare their revision values from left to right. If one version string has fewer revisions, treat the missing revision values as `0`.

Return the following:

- If `version1 < version2`, return `-1`.
- If `version1 > version2`, return `1`.
- Otherwise, return `0`.

## Examples

### Example 1

```text
Input: version1 = "1.2", version2 = "1.10"
Output: -1
```

**Explanation:**

`version1`'s second revision is `"2"` and `version2`'s second revision is `"10"`. Since `2 < 10`, `version1 < version2`.

### Example 2

```text
Input: version1 = "1.01", version2 = "1.001"
Output: 0
```

**Explanation:**

Ignoring leading zeroes, both `"01"` and `"001"` represent the same integer `1`.

### Example 3

```text
Input: version1 = "1.0", version2 = "1.0.0.0"
Output: 0
```

**Explanation:**

`version1` has fewer revisions, so every missing revision is treated as `0`.

## Constraints

- `1 <= version1.length, version2.length <= 500`
- `version1` and `version2` only contain digits and `.`.
- `version1` and `version2` are valid version numbers.
- All given revisions in `version1` and `version2` can be stored in a 32-bit integer.
