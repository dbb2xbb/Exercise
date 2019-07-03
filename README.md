## 随机生成不同的随即数，要求空间复杂度为O(n)
### 算法核心思想：
  - 申请指定随机长度N的切片s，
  - i 0->N遍历，将[0,N)的值存于[0,N)的下表中，
  - 令j = rand(N-i)+i,此步骤意为选取切片s中下标为[i,N)之间的任意值。因为此前
    是按照递增顺序将0~N之间的数据存放到切片中的，所以能当前索引i之后的值都是不重复的。
  - a[i] <==> a[j]
  - 最终返回a[i]
### 测试函数功能：  
   - `go test`