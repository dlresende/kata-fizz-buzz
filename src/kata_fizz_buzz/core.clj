(ns kata-fizz-buzz.core
  (:gen-class))

(defn is-divisible-by
  [divident divisor]
  (= (mod divident divisor) 0))

(defn guess
  [number]
  (if (is-divisible-by number (* 3 5))
    "FizzBuzz"
    (if (is-divisible-by number 5)
      "Buzz"
      (if (is-divisible-by number 3)
        "Fizz"
        (str number)))))

(defn from-1-to
  [number]
  (range 1 (inc  number)))

(defn fizz-buzz
  [last]
  (map guess (from-1-to last)))

(defn parse-number
  [number]
  (. Long parseLong number))

(defn print-output
  [last]
  (println (fizz-buzz last)))

(defn -main
  [& args]
  (let [last (parse-number (first args))]
    (print
      (reduce
        str
        (interpose
          "\n"
          (fizz-buzz last))))))

