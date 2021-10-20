(print "COUNT DOWN:")
(setf i 11)

(while (setf i (- i 1)) 
    (print "Loop #"i)
)

(print "COUNT UP:")
(setf i 0)
(while (block (setf i (+ i 1)) (<= i 10)) 
    (print "Loop #"i)
)
