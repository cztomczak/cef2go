for i in range(1, 100):
    i2 = i
    if i2 < 10:
        i2 = str("0")+str(i)
    i2 = str(i2)
    name = "tag"+i2+".go"
    with open(name, "w") as f:
        f.write("// Package tag.\npackage tag\n\nprintln(\"tag%s\")\n" % i2)
