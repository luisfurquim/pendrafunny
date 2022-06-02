package pendrafusion

func SetTextDecoration(old string, new byte) string {
   var i int
   for ; i<len(old); i++ {
      if old[i] == new {
         break
      }
   }

   if i==len(old) {
      return old + string([]byte{new})
   }

   return old
}

