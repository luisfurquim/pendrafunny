package pendrafusion

func Handler_Strike(cnv *Converter, sel Node) {
   var f *Format

   f = &(cnv.Format[len(cnv.Format)-1])
   f.TextDecoration = SetTextDecoration(f.TextDecoration, 'S')
   cnv.Apply()

   cnv.convert(sel)
}

