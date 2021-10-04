module.exports = {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "references-empty": [1, "never"],
    "scope-case": [2, "always", "lower-case"],
    "subject-case": [0],
    "type-case": [2, "always", "lower-case"]
  },
}
