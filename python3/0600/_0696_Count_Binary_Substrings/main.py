class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        n, res, prev, cur = len(s), 0, 0, 1
        for i in range(1, n):
            if s[i] == s[i - 1]:
                cur += 1
            else:
                res += min(prev, cur)
                prev = cur
                cur = 1
        return res + min(prev, cur)
