# Đề bài:
Cho một số nguyên x, trả về true nếu x là palindrome, trả về false nếu là số khác

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

# Giải thuật:
Kiểm tra nếu x < 0 loại đc case 2
Tạo ra 1 biến reversed bằng 0

for x = 121 (thỏa mãn x > 0)
reversed = 0 + 1
x = 12

for x = 12 (thỏa mãn x > 0)
reversed = 1*10 + 2 = 12
x = 1

for x = 1 (thỏa mãn x > 0)
reversed = 12*10 + 1 = 121
x = 0

Dừng vòng lặp

Kiểm tra original có bằng reversed không nếu có trả về true



