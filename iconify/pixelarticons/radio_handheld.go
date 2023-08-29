package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioHandheld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2v5h8v15H7V2h2zm0 7v4h6V9H9zm6 6H9v5h6v-5z"/>`),
		g.Group(children),
	)
}