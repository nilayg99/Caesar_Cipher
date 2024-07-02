(* Define the mutable reference for the character set *)
let char_values =  "~`!@#$%^&*()+_-={}|[]:;,./<>?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890";;

(* Define the mutable reference for the shift value *)
let shift = ref 3;;

(* Define the main function *)
let rec main () = 
  Printf.printf "char_values: %s\n" ^ char_values;
  Printf.printf "shift: %d\n" !shift;;

(* Invoke the main function *)
let () = main ();;
