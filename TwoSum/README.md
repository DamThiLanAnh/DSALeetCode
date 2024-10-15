# Đề bài
Cho một mảng các số nguyên nums và một số nguyên target, 
trả về chỉ số của hai số sao cho tổng của chúng bằng target .

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

# Giải thuật
Case1: [2, 7, 11, 15]
vòng for đầu tiên: i = 0; i < 4; i++ (thỏa mãn)
vòng for thứ hai: j = 0+1; j < 4; j++ (thỏa mãn)
câu lệnh điều kiện: if nums[0] + nums[1] == 9 (2+7=9) (thỏa mãn)
trả về 1 mảng truyền vào i, j: [0,1]

Case2: [3, 2, 4] 
vòng for đầu tiên: i = 0; i < 3; i++ (thỏa mãn)
vòng for thứ 2: j = 1; j < 3; j ++ (thỏa mãn)
câu lệnh điều kiện: if nums[0] + nums[1] == 6 (3+2=5) (ko thỏa mã)

Đi tới vòng for đầu tiên: i = 1; i < 3; i++ (thỏa mãn)
Vòng for tiếp theo: j = 2; j < 3; j++ (thỏa mãn)
Câu lệnh điều kiện: if nums[1]+nums[2] == 6 (2+4=6) (t/m)
trả về 1 mảng truyền vào i, j: [1,2]

Case3: [3, 3]
vòng for đầu tiên: i = 0; i < 2; i++ (thỏa mãn)
Vòng for thứ 2: j = 1; j < 2; i++ (thỏa mãn)
Câu lệnh điều kiện: if nums[0] + nums[1] == 6 (3+3=6) (t/m)
Trả về 1 mảng truyền vào i, j: [0,1]

# Kết luận
Độ phức tạp về thời gian: O(n^2)
Độ phức tạp về không gian: O(1)
