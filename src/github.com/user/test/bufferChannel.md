缓冲类型的channel,顾名思义，就是指channel可以设置缓冲区大小,可以设置这个channel的缓冲区可以接收多少个元素。
在没有达到这个缓冲区元素个数上限之前都不会被阻塞。如果channel的缓冲区已经满了，则后续再想缓冲区中写入元素的操作就会被阻塞，
直到其他goroutine从当前channel中读走数据释放缓冲区后，阻塞才会被解除。 对应属性JAVA的人来说，这就很像BlockingQueue.

ch := make(chan type, value)

当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。
