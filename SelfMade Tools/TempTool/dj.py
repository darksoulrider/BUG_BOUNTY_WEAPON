def square(n): 
    result = 0
    for p in range (n) : 
        for q in range (n) : 
            print(f"{p} + {q} => {p+q}",end=" ")
            result +=1
            print(f"Result: {result}")
    return result 
print("----------------")
print(square(5))