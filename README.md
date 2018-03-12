# Multiple-choice cheater simulator

Imagine that you have a multiple-choice exam, which you have no idea how to
answer. However, you know the exact distribution of correct answers. Which
strategy should you choose? This code tries a few options.

Here is a sample run:

```
Using distribution [25 25 25 25]
"RandomExam": {AverageGrade:25.01224 PassProbability:0}
"GuessCommon": {AverageGrade:25 PassProbability:0}
"EliminateAndScale": {AverageGrade:24.99465 PassProbability:0}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [40 25 25 10]
"RandomExam": {AverageGrade:29.49747 PassProbability:0}
"GuessCommon": {AverageGrade:40 PassProbability:0}
"EliminateAndScale": {AverageGrade:31.8891 PassProbability:0}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [50 25 25 0]
"RandomExam": {AverageGrade:37.51658 PassProbability:0.00015}
"GuessCommon": {AverageGrade:50 PassProbability:0}
"EliminateAndScale": {AverageGrade:37.49175 PassProbability:0.00018}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [54 25 21 0]
"RandomExam": {AverageGrade:39.82456 PassProbability:0.00075}
"GuessCommon": {AverageGrade:54 PassProbability:0}
"EliminateAndScale": {AverageGrade:39.79816 PassProbability:0.00064}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [54 25 11 10]
"RandomExam": {AverageGrade:37.62779 PassProbability:7e-05}
"GuessCommon": {AverageGrade:54 PassProbability:0}
"EliminateAndScale": {AverageGrade:41.00943 PassProbability:0.00058}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [54 20 13 13]
"RandomExam": {AverageGrade:36.52747 PassProbability:0}
"GuessCommon": {AverageGrade:54 PassProbability:0}
"EliminateAndScale": {AverageGrade:40.77749 PassProbability:0.00024}
"TrueCheater": {AverageGrade:100 PassProbability:1}

Using distribution [100 0 0 0]
"RandomExam": {AverageGrade:100 PassProbability:1}
"GuessCommon": {AverageGrade:100 PassProbability:1}
"EliminateAndScale": {AverageGrade:100 PassProbability:1}
"TrueCheater": {AverageGrade:100 PassProbability:1}
```
