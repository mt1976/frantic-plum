package contentWrangler

type fartFarmer struct {
  name string
  id string
}

var example fartFarmer

func new(in string) fartFarmer {
  var out fartFarmer
  out.name = in
  out.id = in
  return out
}

func init(){
  example = new("example")
}
