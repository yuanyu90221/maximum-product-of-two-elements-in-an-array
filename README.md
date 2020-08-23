# maximum-product-of-two-elements-in-an-array

## 題目解讀：

### 題目來源:
[maximum-product-of-two-elements-in-an-array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/)

### 原文:
Given the array of integers nums, you will choose two different indices i and j of that array. Return the maximum value of (nums[i]-1)*(nums[j]-1).

### 解讀：

給定一個正整數陣列nums, 任意選兩個index i,j 求出 最大值的 (nums[i]-1)* (nums[j]-1)


## 初步解法:
### 初步觀察:
首先 因為都是正整數

所以知道 任意正整數 a, b,  a * b的值跟a,b值為正相關

a,b 值越大 則 a * b 值越大

所以 這題相當於是在找尋 陣列最大的兩個值


### 初步設計:
Given a integer array nums

Step 0: let i = 0, max = 0, secondMax = 0

Step 1: if i >= len(nums) go to step 4

Step 2: if nums[i] > max  then  set secondMax = max,  max = nums[i]

Step 3: if nums[i] <= max && nums[i] > secondMax  set secondMax = nums[i]

Step 4: return (max - 1)* (secondMax - 1)
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package max_product

func maxProduct(nums []int) int {
	result := 0
	max := 0
	secondMax := 0
	for _, val := range nums {
		if val > max {
			secondMax = max
			max = val
		} else if val > secondMax {
			secondMax = val
		}
	}
	result = (max - 1) * (secondMax - 1)
	return result
}

```
## 測資的撰寫

```golang
package max_product

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{3, 4, 5, 2},
			},
			want: 12,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 5, 4, 5},
			},
			want: 16,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{3, 7},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 24th](https://hackmd.io/BKK9QlHKTcq9agHLc3zrRw?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)