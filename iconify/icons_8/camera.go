package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.5 6l-.313.406L10 8H3v18h26V8h-7l-1.188-1.594L20.5 6h-9zm1 2h7l1.188 1.594L21 10h6v14H5V10h6l.313-.406L12.5 8zM8 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm8 0c-3.302 0-6 2.698-6 6s2.698 6 6 6s6-2.698 6-6s-2.698-6-6-6zm0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4z"/>`),
		g.Group(children),
	)
}