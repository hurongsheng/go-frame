###Map [底层原理](https://www.jianshu.com/p/26f9be26509e)

#### map的数据结构

    map的底层实现是个HashTable，并且使用链表解决哈希冲突
    数据结构是hmap，是hashmap的缩写
    1.发生hash冲突时会扩容
    2.key和value是各自放一起的(某些情况下可以减少padding字段，节省空间)
    3.2^B=bmap的数量，每个 bmap可以存8份数据，第9份数据会以overflow的链表连到新的bmap上
    
#### [map如何扩容](https://www.jianshu.com/p/26f9be26509e)

    碰撞会产生overflow 链表。假设所有数据都是碰撞的，那么hash表就没有价值了，等价于链表的O(n)检索。
    所有会有一个装载因子，用于表示总的bmap数量和当前的总数量的占比，源码中是6.5
    扩容有两种，
        - 一种桶装满了的，B+1,也就是桶数量翻倍，此时保持序号不变即可。
        - 另一种是overflow桶太多，但是没怎么满，需要rehash
    扩容是渐进式的
    
#### map 中的 key 为什么是无序的？
    
    虽然是maphash table，但也维护了buckets，里面存的是bmap的数组指针。按理说是遍历应该有序的。
    但是无序是因为map遍历的时候增加了一个随机数（主要目的是设计者为了避免新手误以为map是有序的引发问题）
    扩容后buckets有可能改顺序
    
   
#### map 是线程安全的吗

    no，读写不安全，slice也是
    为什么不安全
           - 扩容不是原子的
           - 保存buckets的是slice，不是线程安全的