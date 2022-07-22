package forcast

type IRepository interface {
  Create(string) Forcast
  Get(int, string) Forcast
  List() []Forcast
}
