package organization

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintOut(list []Organization) {
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "Organizatiopn Name", "ID", "Company ID")
	//	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "", "", "")

	for _, v := range list {
		fmt.Fprintf(w, "\n %v\t%v\t%v\t", v.Name, v.ID, v.CompanyID)
	}

}
