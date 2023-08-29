package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontbase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.079 13.996c-2.702-2.771-5.702-5.703-8.105-8.103c-1.62-1.621-4.284-1.621-5.943 0c-2.97 2.963-5.248 5.21-8.104 8.066a3.12 3.12 0 0 0 0 4.437a3.12 3.12 0 0 0 4.437 0l2.2-2.2l2.2 2.2a3.12 3.12 0 0 0 4.438 0a3.12 3.12 0 0 0 0-4.438l4.4 4.4a3.12 3.12 0 0 0 4.438 0c1.274-1.16 1.274-3.165.039-4.362z"/>`),
		g.Group(children),
	)
}