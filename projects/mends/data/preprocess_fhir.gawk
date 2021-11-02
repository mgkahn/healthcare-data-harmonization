BEGIN {
 ORS=""
 if (tolower(ARGV[1])~/condition/) prefix = "Condition_Occurrence"
 if (tolower(ARGV[1])~/drug/) prefix = "Drug_Exposure"
 if (tolower(ARGV[1])~/measurement/) prefix = "Measurement" 
 if (tolower(ARGV[1])~/provider/) prefix = "Provider"
 pre_json = "{ \"" prefix "\" : \n[\n"  
 print(pre_json)
}
{ if (NR > 1 && NR<=5) {printf(", \n")} 
  if (NR <= 5) { print $0 } 
}
END {post_json = "\n]\n}\n" 
     print(post_json)
}


