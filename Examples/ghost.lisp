(setf choiceText "There are 3 doors, a ghost behind 1, which door do you choose: ")
(setf alive 1)
(setf score 0)

(while (== alive 1) (block 
    (setf choice (intin choiceText))
    (setf ghost (+ (rnd 3) 1))
    (if (== choice ghost) 
        (setf alive 0)
        (block (print "You Survived") (setf score (+ score 1)))
    )
))
(print "You died, your score was " score)
