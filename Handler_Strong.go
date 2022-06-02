package pendrafusion

func Handler_Strong(cnv *Converter, sel Node) {
   var f *Format

   f = &(cnv.Format[len(cnv.Format)-1])
   f.TextDecoration = SetTextDecoration(f.TextDecoration, 'B')
   cnv.Apply()

   cnv.convert(sel)
}

