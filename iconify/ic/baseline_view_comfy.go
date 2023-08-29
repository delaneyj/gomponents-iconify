package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineViewComfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v7h20V4H2zm8 16h12v-7H10v7zm-8 0h6v-7H2v7z"/>`),
		g.Group(children),
	)
}