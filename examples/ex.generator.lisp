; Lazy generator example

(def take (n lst) (take n lst '()))
(def take (0 lst acc) acc)
(def take (n lst acc) (take (- n 1) (tail lst) (append acc (head lst))))

; iterator: state -> (value, new-state)
(def from-int (state) (list state (+ state 1)))

; gen: iterator initial-state -> lazy-list
(set ints (gen from-int 1))

(print (take 20 ints))


; iterator: state -> (new-state)
(def next-int (n) (list (+ n 1)))
(set ints2 (gen next-int 0))

(print (take 20 ints2))


; finite iterator: state -> '()
(def next-below-10 (10) '())
(def next-below-10 (n) (list (+ n 1)))
(set ints3 (gen next-below-10 0))

; print evaluates lazy list
(print ints3)


; iterator: state -> new-state
(def inc (n) (+ n 1))
(set ints4 (gen inc 0))
(print (take 15 ints4))


; iterator: state ... -> value, new-state ...
(def next-fib (a b) (list b b (+ a b)))
(set fibs (gen next-fib 1 1))
(print (take 10 fibs))

; vim: ft=lisp
