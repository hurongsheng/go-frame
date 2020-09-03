
### go指针和unsafe.point和 unitptr的区别

    1.go不同数据类型的指针不能相互赋值，不能比较，不能数学运算
    2.任何类型的指针都可以跟unsafe.point做类型转换
    3.unitptr就可以跟unsafe.point相互转化
    4.unitptr没有指针语义，unitptr指向的数据会被gc回收

