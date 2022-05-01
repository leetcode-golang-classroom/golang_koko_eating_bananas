# golang_koko_eating_bananas

Koko loves to eat bananas. There are `n` piles of bananas, the `ith` pile has `piles[i]` bananas. The guards have gone and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of `k`. Each hour, she chooses some pile of bananas and eats `k` bananas from that pile. If the pile has less than `k` bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return *the minimum integer* `k` *such that she can eat all the bananas within* `h` *hours*.

## Examples

**Example 1:**

```
Input: piles = [3,6,7,11], h = 8
Output: 4

```

**Example 2:**

```
Input: piles = [30,11,23,4,20], h = 5
Output: 30

```

**Example 3:**

```
Input: piles = [30,11,23,4,20], h = 6
Output: 23
```

**Constraints:**

1 ≤ piles.length ≤ $10^4$
piles.length ≤ h ≤ $10^9$
1 ≤ piles[i] ≤ $10^9$

## 解析

題目給定一個整數陣列 piles 來表示香蕉放在每個 pile的個數，piles[i] 代表第 i 個 pile 所裝的香蕉 個數。給定一個時間 h 小時。假設 koko 吃每小時吃 k 根香蕉 ，求 k 最小值要是多少才能夠吃完所有 piles 內的香蕉。

另外每小時只能在一個 pile 內吃香蕉

思考我們怎麼去計算每個房間吃香蕉花的時間

假設每個小時吃 k 根 bananas

那再一間房間所花的時間就是 (piles[i]/k) + (piles[i]%k ≠ 0 ? 1:0) 

把上面的那算式累加起來就是所有房間花的時間

要找到最小時間 可以透過 k= 1 開始計算看看所有花費時間是不是有 ≤ h

這樣去找總共要花 O(n*m) 其中 m 是那個最小香蕉個數

然後透過以下兩個特性可以將演算法優化

1. 如果每小時吃了 n 根可以達成 —> n+1 根也可以
2. 如果每小時吃了 n 根無法達成 —> n-1 根也無法

因為每小時只能在一個 pile 內吃香蕉， 所以找出一個pile裝最多香蕉個數就是達成條件某個上限

只要能在一個小時內吃完含最多香蕉的pile 其他pile 也一定可以在一小時內做到

所以可以對香蕉數目做 binary search

把 min := 1, max := max(piles)

每次計算 mid = (min+max)/2

然後透過累加  hspent = (piles[i]/k) + (piles[i]%k ≠ 0 ? 1:0)

當 hspent > h 時, 把 min = mid + 1 

當 hspent ≤ h 時, 把 max = mid

當 min == max ，max 就是答案

這樣每次搜詢 O(n*logm) 其中 m 是 piles 中的最大值

## 程式碼

```go
func minEatingSpeed(piles []int, h int) int {
    min := 1
    max := 0
    for _, val := range piles {
        if max < val {
            max = val
        }
    }
    for min < max  { 
        mid := (min+max)/ 2
        hSpent := 0
        for _, val := range piles {
            hSpent += val/mid
            if val % mid != 0 {
                hSpent +=1
            }
        }
        if hSpent <= h {
            max = mid
        } else {
            min = mid + 1
        }
    }
    return max
}
```

## 困難點

1. 需要理解英文敘述裡第2段所說 每小時只能在一個 piles 吃香蕉的條件
2. 需要知道怎麼去算花費時間
3. 更新條件的算法不同於一般的 binary search

## Solve Point

- [x]  Understand what problem want to solve
- [x]  Figure which Data structure is appropriate
- [x]  Find the key condition for Binary Search