(setf num (+ (rnd 10) 1))
(setf result (if (>= num 5) 
    (if (== num 5)
        "THE SAME AS"
        "LARGER THAN"
    )
    "SMALLER THAN"
))
(print num " IS " result " 5")

