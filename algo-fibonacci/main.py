import time

def Fibonacci(n):
    if n < 2:
        return n
    
    return Fibonacci(n-1) + Fibonacci(n-2)

start = time.time()
Fibonacci(25)
print("Excution time", time.time() - start)