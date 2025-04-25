
# 🧠 Golang Function Types Explained in Bengali (with Examples)

এই ডকুমেন্টে আমরা Golang এর বিভিন্ন ধরনের ফাংশনের ব্যাখ্যা, উপমা, এবং উদাহরণ সহ দেখব।

---

## ✅ ১. Standard or Named Function

📘 **সংজ্ঞা**:  
যেসব ফাংশনের নির্দিষ্ট নাম থাকে এবং এগুলো সাধারণভাবে func কীওয়ার্ড ব্যবহার করে ডিফাইন করা হয়, সেগুলোকে standard বা named function বলা হয়।

🧠 **ব্যাখ্যা**:  
এটি ঠিক যেন আমরা কাউকে ফোন নাম্বার দিয়ে কল করি — যেমন, `Call("Mom")`।  

✅ **উদাহরণ**:
```go
package main
import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    msg := greet("Shakib")
    fmt.Println(msg) // Output: Hello, Shakib
}
```

---

## ✅ ২. Anonymous Function

📘 **সংজ্ঞা**:  
যে ফাংশনের কোনো নাম নেই, সেটি anonymous function। সাধারণত এদের কোনো ভেরিয়েবলে অ্যাসাইন করে ব্যবহার করা হয়।

🧠 **ব্যাখ্যা**:  
একটি ক্ষণস্থায়ী helper – নাম নেই, কাজ শেষ করে চলে যায়।

✅ **উদাহরণ**:
```go
func() {
    fmt.Println("Anonymous Function")
}()
```

```go
sayHi := func(name string) {
    fmt.Println("Hi", name)
}
sayHi("Shakib")
```

---

## ✅ ৩. Function Expression / Assign Function in Variable

📘 **সংজ্ঞা**:  
একটি ফাংশনকে কোনো ভেরিয়েবলে সংরক্ষণ করা এবং পরে ব্যবহার করা হলে তাকে function expression বলা হয়।

🧠 **ব্যাখ্যা**:  
এটি হচ্ছে ফাংশনকে gift-wrap করে রাখা — চাইলে খুলে ব্যবহার করো।

✅ **উদাহরণ**:
```go
add := func(a int, b int) int {
    return a + b
}
fmt.Println(add(5, 3)) // Output: 8
```

---

## ✅ ৪. Higher Order Function / First-Class Function

📘 **সংজ্ঞা**:  
যেসব ফাংশন অন্য ফাংশনকে প্যারামিটার হিসেবে নিতে পারে অথবা রিটার্ন করতে পারে, তাদের Higher-order বা First-class function বলে।

🧠 **ব্যাখ্যা**:  
একটি ফাংশন যেটা আরেকটি ফাংশনকে ট্রেইনার হিসেবে ডাকে।

✅ **উদাহরণ**:
```go
func doOperation(a int, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    multiply := func(x int, y int) int {
        return x * y
    }
    fmt.Println(doOperation(4, 5, multiply)) // Output: 20
}
```

---

## ✅ ৫. Callback Function

📘 **সংজ্ঞা**:  
যেসব ফাংশনকে আরেকটি ফাংশনে প্যারামিটার হিসেবে পাঠানো হয় এবং সেই ফাংশনের ভিতরে ডাকা হয়, তাদের callback function বলা হয়।

🧠 **ব্যাখ্যা**:  
তুমি কাউকে বলো, "তুমি কাজ শেষ হলে আমাকে কল করো" – ঐ কলটাই callback।

✅ **উদাহরণ**:
```go
func process(name string, callback func(string)) {
    fmt.Println("Processing...")
    callback(name)
}

func main() {
    process("Shakib", func(n string) {
        fmt.Println("Welcome", n)
    })
}
```

---

## ✅ ৬. Variadic Function

📘 **সংজ্ঞা**:  
যেসব ফাংশন একাধিক সংখ্যক আর্গুমেন্ট নিতে পারে, সেগুলো variadic function।

🧠 **ব্যাখ্যা**:  
তুমি বলছো – "আমি যত খুশি সংখ্যা দিব, তুমি সব যোগ করে দাও"।

✅ **উদাহরণ**:
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // Output: 6
    fmt.Println(sum(4, 5, 6, 7, 8)) // Output: 30
}
```

---

## ✅ ৭. Init Function

📘 **সংজ্ঞা**:  
Go তে একটি বিশেষ ফাংশন `init()` আছে যেটি main function চলার আগে একবার execute হয়।

🧠 **ব্যাখ্যা**:  
তুমি ঘুম থেকে ওঠার পরে দাঁত ব্রাশ করো — `init()` হচ্ছে সেই প্রথম কাজ।

✅ **উদাহরণ**:
```go
func init() {
    fmt.Println("Init: App is starting...")
}

func main() {
    fmt.Println("Main: Hello world")
}
```

📤 Output:
```
Init: App is starting...
Main: Hello world
```

---

## ✅ ৮. Closure

📘 **সংজ্ঞা**:  
যেসব ফাংশন বাইরের স্কোপের ভেরিয়েবলকে মনে রাখতে পারে এবং ব্যবহার করতে পারে, তাদের closure বলা হয়।

🧠 **ব্যাখ্যা**:  
একটা ফাংশনের স্মৃতিশক্তি আছে – আগের ভ্যালু মনে রাখে।

✅ **উদাহরণ**:
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    next := counter()
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
}
```

---

## ✅ ৯. Defer Function

📘 **সংজ্ঞা**:  
`defer` কীওয়ার্ড দিয়ে যেসব ফাংশন কল করা হয়, সেগুলো function শেষ হওয়ার পরে এক্সিকিউট হয় (যেমন `finally` ব্লকের মতো)।

🧠 **ব্যাখ্যা**:  
তুমি বলো, "শেষে দরজা বন্ধ করে দিও" – দরজা বন্ধ defer করা।

✅ **উদাহরণ**:
```go
func main() {
    defer fmt.Println("Cleaning up")
    fmt.Println("Doing work")
}
```

📤 Output:
```
Doing work
Cleaning up
```

---

## ✅ ১০. Receiver Function (Method Function)

📘 **সংজ্ঞা**:  
যখন কোনো struct টাইপের উপর method তৈরি করা হয়, তখন তাকে receiver function বলা হয়।

🧠 **ব্যাখ্যা**:  
এটি একটি struct-এর member function এর মতো।

✅ **উদাহরণ**:
```go
type Person struct {
    name string
}

func (p Person) greet() {
    fmt.Println("Hello,", p.name)
}

func main() {
    p := Person{name: "Shakib"}
    p.greet()
}
```

---

## ✅ ১১. IIFE (Immediately Invoked Function Expression)

📘 **সংজ্ঞা**:  
যে ফাংশনটি ডিফাইন করার সাথে সাথে একবার কল করে ফেলা হয়, তাকে IIFE বলা হয়।

🧠 **ব্যাখ্যা**:  
তুমি বলো: “আমি এখনই বলে দিচ্ছি” – সঙ্গে সঙ্গে বলা হয়ে গেল।

✅ **উদাহরণ**:
```go
func main() {
    func(msg string) {
        fmt.Println("Message:", msg)
    }("Go is awesome!")
}
```

---

## 🎁 সংক্ষেপে মনে রাখার ট্রিক:

| ফাংশনের ধরন | কাজ |
|-------------|-----|
| Named | নাম দিয়ে ডাকা যায় |
| Anonymous | নামহীন, কিন্তু কাজ করে |
| Expression | ভ্যারিয়েবলে রাখা |
| Higher Order | অন্য ফাংশন নেয় বা রিটার্ন করে |
| Callback | ডাকা হয় পরে |
| Variadic | অনেক ইনপুট নেয় |
| Init | শুরুতে চলে |
| Closure | মেমোরি রাখে |
| Defer | শেষে চলে |
| Receiver | Struct এর মেথড |
| IIFE | বানিয়ে সাথে সাথে চলে |
