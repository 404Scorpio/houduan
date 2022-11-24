package csv


func WriteCSV(c[][]string){

	f, err := os.Open("abc.csv",0644)
	if nil != err {
    fmt.Println("open", abc.csv, "failed!")
    return
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	//写数据
	f, _ := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0666)
	defer f.Close()
	_, err = f.WriteString("target_compile_option")
	if nil != err {
		fmt.Println(err)
	}

	w.WriteAll()
	w.Flush()
}

func ReadCSV() [][]string {
	fileName := "abc.csv"
	fs1, _ := os.Open(abc.csv)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

	return content
}

func ShowCSV() {
	fileName := "Debt.csv"
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content
}