# What will I learn 
- Strongly Typed 
- Complied Languages
  - Anything you have to build before you run
- At bigger companies using multiple languages because different languages are efficient for different things

What will I have done
- Built a slack bot

Q's
- How does someone best succeed in this course
  # Really looking for progress overtime
  - So you and your instructor can tell who you are as an engineer

## Final
- Repl.it Code Exam

## Resources to Learn Fast
- golang docs
- exercism

## When you make something
- Make it is easy for me to use your thing fast

## Syntics
- main is the first package you create (strongly typed entry point)

### Go is phonoemally close to the hardware
- Swift goes through software to run the code
- Go is more explicit 
  - import everything you need

# Work
2. Watch Go tools screencast (5 min)
   - Watch Screencast
3. Complete the Tour of Go on golang.org
  - Get comfortable with playground and go syntic
  - Show Dani when done
4. Start setting up Exercism
  - Write Go on your own
5. Work on tutorials

- Editor Plugin and IDE

Where I left off on the Tour of Go tutorial: https://tour.golang.org/moretypes/15

Another resource/video I might watch:
https://www.youtube.com/watch?v=29LLRKIL_TI

### Tour of Go
- Exercise Loops and Functions
    ```go
    package main

    import (
      "fmt"
    )

    func Sqrt(x float64) (float64, int) {
      z := float64(x/2)
      count := 1
      prev_z := float64(0)
      for prev_z != z {
        prev_z = z
        z -= (z*z - x) / (2*z)
        fmt.Println(z)
        count += 1
      }
      return z, count
    }

    func main() {
      fmt.Println(Sqrt(4))
    }
    ```