(setf i 0)
(while (< i 10) (block 
    (print i)
    (setf i (+ i 1))
))

(print "DONE")
