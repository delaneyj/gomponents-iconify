package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 3.5h8v8m-8-8V0m0 3.5H0m11.5 8h-8V6m8 5.5V15m0-3.5H15"/>`),
		g.Group(children),
	)
}