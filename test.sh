go run main.go tests/sample1.txt res1.txt
go run main.go tests/sample2.txt res2.txt
go run main.go tests/sample3.txt res3.txt
go run main.go tests/sample4.txt res4.txt
go run main.go tests/sample5.txt res5.txt
go run main.go tests/sample6.txt res6.txt
go run main.go tests/sample7.txt res7.txt
cat tests/sample1.txt; echo -e; cat tests/result1.txt; echo -e; cat res1.txt; echo -e;
read w <tests/result1.txt; read r < res1.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e; echo -e
cat tests/sample2.txt; echo -e; cat tests/result2.txt; echo -e; cat res2.txt; echo -e;
read w <tests/result2.txt; read r < res2.txt;
if [[ "$w" == "$r" ]]; then  echo -e OK
else echo -e "not OK"
fi
echo -e; echo -e
cat tests/sample3.txt; echo -e; cat tests/result3.txt; echo -e; cat res3.txt; echo -e;
read w <tests/result3.txt; read r < res3.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e; echo -e
cat tests/sample4.txt; echo -e; cat tests/result4.txt; echo -e; cat res4.txt; echo -e;
read w <tests/result4.txt; read r < res4.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e; echo -e
cat tests/sample5.txt; echo -e; cat tests/result5.txt; echo -e; cat res5.txt; echo -e;
read w <tests/result5.txt; read r < res5.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e; echo -e
cat tests/sample6.txt; echo -e; cat tests/result6.txt; echo -e; cat res6.txt; echo -e;
read w <tests/result6.txt; read r < res6.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e; echo -e
cat tests/sample7.txt; echo -e; cat tests/result7.txt; echo -e; cat res7.txt; echo -e;
read w <tests/result7.txt; read r < res7.txt;
if [[ "$w" == "$r" ]]; then
    echo -e OK;
else 
    echo -e "not OK";
fi
echo -e