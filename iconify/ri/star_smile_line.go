package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSmileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.999.5l4.226 6.183l7.186 2.109l-4.575 5.93l.216 7.486l-7.053-2.518l-7.054 2.518l.216-7.486l-4.575-5.93l7.187-2.109L11.999.5Zm0 3.544L9.02 8.402L3.956 9.887l3.225 4.178l-.153 5.275l4.97-1.774l4.97 1.774l-.151-5.274l3.224-4.179l-5.065-1.485L12 4.044Zm-2 7.956a2 2 0 1 0 4 0h2a4 4 0 1 1-8 0h2Z"/>`),
		g.Group(children),
	)
}