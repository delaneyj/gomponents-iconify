package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterNegativeNumber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.5 22.5v-1h3v-11h-3v-1h4v12h3v1h-7z"/><path fill="currentColor" d="M23 10v12v-12m1-1h-5v2h3v10h-3v2h8v-2h-3V9zM5 15h10v2H5z"/>`),
		g.Group(children),
	)
}