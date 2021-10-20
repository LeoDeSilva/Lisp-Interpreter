(defun compare (num1 num2) 
  (if (== num1 num2) (block
      (print "Are Equal")
      1
    ) 

    (block 
      (print "Are not equal")
      0
    )
  )
)

(compare 1 2)
