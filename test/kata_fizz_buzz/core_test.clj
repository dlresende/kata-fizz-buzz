(ns kata-fizz-buzz.core-test
  (:require [clojure.test :refer :all]
            [kata-fizz-buzz.core :refer :all]))

(deftest guess-test
  (testing "should print the number it received as argument"
    (is (=  (guess 1) "1")))

  (testing "should print Fizz when a number is divisible by 3"
    (is (= (guess 6) "Fizz")))

  (testing "should print Buzz when a number is divisible by 5"
    (is (= (guess 10) "Buzz")))

  (testing "should print FizzBuzz when a number is divisible by 3 and 5"
    (is (= (guess 15) "FizzBuzz"))))

(deftest main-test
  (testing "should play fizz-buzz until 5"
    (is (= (fizz-buzz 5) '("1", "2", "Fizz", "4", "Buzz")))))
