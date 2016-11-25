package stringutil

func FullName(f, l string) (string, int) {
  full := f + " " + l
  length := len(full)
  return full, length
}

func FullNameNakedReturn(f, l string) (fn string, cc int) {
  fn = f + " " + l
  cc = len(fn)
  return
}
