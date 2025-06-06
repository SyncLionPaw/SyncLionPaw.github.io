# 系统调用

## 系统调用的流程
以xv6为例，理解系统调用的流程

1. 用户程序使用库函数，例如glibc
2. 库函数中使用了系统函数，例如write，read，也要将参数放到a0-a5。
3. 系统函数内部，将系统调用号保存到某个寄存器（a7），呼叫ecall
4. ecall做以下事情
    0. 保存用户的通用寄存器，上下文，便于后面恢复
    1. 保存当前程序指针 pc 到sepc，保存当前特权级别到sstatus的特定位，写scause，中断原因
    2. 切换到内核态，关中断
    3. 读取stvec, 中断向量表，找到下一步要去的地方 uservec
    4. 读取系统调用号码a7和参数a0-a5
    5. 做对应的系统调用操作
    6. 完成后恢复用户上下文
    7. 开中断
    8. sret 将sepc恢复到pc
5. 恢复到用户程序，判断系统调用的返回结果

### 用户程序使用库函数

可以使用 strce 看自己的程序里面使用了哪些系统调用。

### 库函数内部使用系统调用

## 详细代码流程

### ecall，stvec，sepc

ecall是用户级别的指令，不是特权指令，用户态可以执行这个指令。
sepc保存当前的PC后面再恢复的。

### stvec
stvec就是中断向量表寄存器。保存了：发生trap的时候，要去哪里执行中断处理函数。

这个寄存器用最低位保存模式，它有两种模式。

在riscv中，这个stvec保存的是uservec这段汇编指令的地址

### uservc
uservcui汇编指令会完成以下内容：
1. 保存上下文到进程的trampframe
2. 内核栈，内核页表
3. 跳转到 usertrap

### usertrap

### usertrapret


## 系统调用的实现

### 和文件系统打交道
