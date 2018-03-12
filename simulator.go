// Multiple-choice cheater simulator
//
// Imagine that you have a multiple-choice exam, which you have no idea how to
// answer. However, you know the exact distribution of correct answers. Which
// strategy should you choose? This code tries a few options.

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	ExamLength         = 100
	OptionsPerQuestion = 4
	PassGrade          = 55
	NumIterations      = 100000
)

// Answer is an answer to a multiple-choice exam. Between 0 and
// OptionsPerQuestion.
type Answer int

// Exam is the set of responses to a multiple-choice exam. It can represent
// either the correct answers for an exam or an attempt to answer an exam.
type Exam [ExamLength]Answer

// Grade is the grade for an exam, in the range 0-ExamLength.
type Grade int

// Distribution is the distribution of answers for an Exam.  For a
// given distribution d, the number of "0" answers in an exam should
// be d[0], the number of "1" answers should be d[1], etc.
type Distribution [OptionsPerQuestion]int

// Check returns the grade that answers should get if the
// correct answers are e.
func (e *Exam) Check(answers *Exam) Grade {
	var grade Grade = 0

	for i := range e {
		if e[i] == answers[i] {
			grade++
		}
	}

	return grade
}

// Distribution returns the Distribution of answers in e.
func (e *Exam) Distribution() Distribution {
	var d Distribution

	for _, a := range e {
		d[a]++
	}

	return d
}

// Valid returns true iff the sum of the distribution d is ExamLength.
func (d Distribution) Valid() bool {
	sum := 0
	for _, v := range d {
		if v < 0 {
			return false
		}
		sum += v
	}
	return sum == ExamLength
}

// Sum returns the sum of a distribution
func (d Distribution) Sum() int {
	sum := 0
	for _, v := range d {
		sum += v
	}
	return sum
}

// Rank returns an array of OptionPerQuestion Answers, with the first being the
// most common one in d, the second being second-most-common, etc.
func Rank(d Distribution) [OptionsPerQuestion]Answer {
	positions := make([]struct {
		position  Answer
		frequency int
	}, OptionsPerQuestion)

	var result [OptionsPerQuestion]Answer

	for i := range positions {
		positions[i].position = Answer(i)
		positions[i].frequency = d[i]
	}

	sort.Slice(positions, func(i, j int) bool {
		return positions[i].frequency > positions[j].frequency
	})

	for i := range result {
		result[i] = positions[i].position
	}

	return result
}

// Strategy is a strategy for answering a random exam that matches a given distribution.
type Strategy func(Distribution) Exam

// RandomExam returns a random exam with distribution d.
func RandomExam(d Distribution) Exam {
	var tmp, dest Exam

	var i = 0

	for a := Answer(0); a < OptionsPerQuestion; {
		if d[a] == 0 {
			a++
		} else {
			tmp[i] = a
			d[a]--
			i++
		}
	}

	perm := rand.Perm(ExamLength)

	for i, v := range perm {
		dest[v] = tmp[i]
	}

	return dest
}

// GuessCommon is a Strategy which simply guesses the most common answer is the answer for all questions.
func GuessCommon(d Distribution) Exam {
	var e Exam

	common := Rank(d)[0]

	for _, i := range e {
		e[i] = common
	}

	return e
}

// EliminateAndScale is a strategy which eliminates the least-common answer and
// scales the remaining answers up accordingly, and answers the test based on
// that.
func EliminateAndScale(d Distribution) Exam {
	ranking := Rank(d)

	d[ranking[OptionsPerQuestion-1]] = 0

	sum := d.Sum()

	for i := range d {
		d[i] = int(float64(d[i]) * float64(ExamLength) / float64(sum))
	}

	// Eliminate remainder artifacts
	sum = d.Sum()
	d[0] += ExamLength - sum

	return RandomExam(d)
}

// Assessment is an assessment of how well a strategy works
type Assessment struct {
	AverageGrade    float64
	PassProbability float64
}

// AssessStrategy performs multiple iterations of creating a random with
// distribution d, using Strategy s to try and solve it, and assessing the
// result. If s is nil, the strategy is to "truly cheat", giving the actual
// correct answers for the exam (this is for debugging purposes).
func AssessStrategy(s Strategy, d Distribution, iterations int) Assessment {
	passes := 0
	totalGrade := 0

	for i := 0; i < iterations; i++ {
		correct := RandomExam(d)
		var guess Exam
		if s == nil {
			// This is a cheater
			guess = correct
		} else {
			guess = s(d)
		}
		grade := correct.Check(&guess)
		totalGrade += int(grade)
		if grade >= PassGrade {
			passes += 1
		}
	}

	return Assessment{
		AverageGrade:    float64(totalGrade) / float64(iterations),
		PassProbability: float64(passes) / float64(iterations),
	}
}

func main() {
	strategies := []struct {
		name     string
		strategy Strategy
	}{
		{"RandomExam", RandomExam},
		{"GuessCommon", GuessCommon},
		{"EliminateAndScale", EliminateAndScale},
		{"TrueCheater", nil},
	}

	distributions := []Distribution{
		{25, 25, 25, 25},
		{40, 25, 25, 10},
		{50, 25, 25, 0},
		{54, 25, 21, 0},
		{54, 25, 11, 10},
		{54, 20, 13, 13},
		{100, 0, 0, 0},
	}

	for _, d := range distributions {
		fmt.Println("Using distribution", d)

		if !d.Valid() {
			fmt.Println("Invalid distribution:", d)
			continue
		}
		for _, s := range strategies {
			assessment := AssessStrategy(s.strategy, d, NumIterations)
			fmt.Printf("%q: %+v\n", s.name, assessment)
		}
		fmt.Println()
	}
}
