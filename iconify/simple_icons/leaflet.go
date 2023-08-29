package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaflet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.69 0c-.355.574-8.432 4.74-10.856 8.649c-2.424 3.91-3.116 6.988-2.237 9.882c.879 2.893 2.559 2.763 3.516 3.717c.958.954 2.257 2.113 4.332 1.645c2.717-.613 5.335-2.426 6.638-7.508c1.302-5.082.448-9.533-.103-11.99A35.395 35.395 0 0 0 17.69 0zm-.138.858l-9.22 21.585l-.574-.577Z"/>`),
		g.Group(children),
	)
}