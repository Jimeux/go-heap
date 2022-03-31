# go-heap

Functions to get the top K scores from a file.

```
   MIN HEAP
   ✅ parent < child

   parent = (child - 1) / 2
   left   = parent * 2 + 1
   right  = parent * 2 + 2

   0  1  2  3  4
   [1, 2, 3, 5, 4]

       1
      / \
     2   3       // complete tree = insert from left to right on each level
    / \
   5   4        // up
                // down
```
