find . -name "*.sh" | sed 's/.sh//g' | rev | cut -f1 -d"/" | rev
#find . -name "*.sh" | sed 's/.sh//g' |cut -f1 -d"/"
#sed for pattern replacement