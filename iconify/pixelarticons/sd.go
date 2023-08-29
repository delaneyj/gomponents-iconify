package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2h2v20H4V6h2v14h12V4H8V2h10zM8 4H6v2h2V4zm6 2h2v4h-2V6zm-2 0h-2v4h2V6z"/>`),
		g.Group(children),
	)
}