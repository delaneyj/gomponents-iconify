package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cityzen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.77 8v17.23H5.5A17.26 17.26 0 0 1 22.77 8Zm0 0H40a17.27 17.27 0 0 1-17.23 17.23ZM5.5 25.23h17.27V42.5A17.26 17.26 0 0 1 5.5 25.23Zm34.52 0V42.5H22.76A17.27 17.27 0 0 1 40 25.23Zm2.48 0"/>`),
		g.Group(children),
	)
}