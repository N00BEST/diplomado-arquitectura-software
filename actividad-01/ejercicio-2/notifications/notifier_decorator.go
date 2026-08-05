package notifications

// Acá se pudiesen agregar otras funciones en caso
// de ser necesario. Por la simplicidad del ejercicio
// se omiten. Esto debería ser una clase abstracta
// pero Go no maneja clases.
type NotifierDecorator interface {
	Notifier
}
