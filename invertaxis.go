package main;
import (
	"fmt";
	"regexp";
	"os";
	)
func main(){
	if(len(os.Args) <= 2){
		fmt.Println("invertaxis %f %axis");
		os.Exit(0);
	}
	axis := os.Args[2];
	if(strings.Contains(axis,"$")) {
		fmt.Println("no $ in pattern");
		os.Exit(-1);
	}
	if(axis != "x" && axis != "y" && axis != "z"){
		fmt.Println("warning",axis,"not xyz");
	}
	filecontents ,err := os.ReadFile(os.Args[1]);
	if err != nil {
		fmt.Println(err);
		os.Exit(-1);
	}
	gcode := string(filecontents);
	negative := regexp.MustCompile(axis+`(\-)([\d])`); 
	positive := regexp.MustCompile(axis+`([\d])`);
	tmp := regexp.MustCompile(axis+`(\+)([\d])`);

	gcode = positive.ReplaceAllString(gcode,axis+"+$1");
	gcode = negative.ReplaceAllString(gcode,axis+"$2")
	gcode = tmp.ReplaceAllString(gcode,axis+"-$2")
	fmt.Println(gcode);

}
