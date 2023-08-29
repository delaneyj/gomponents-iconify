package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 29H3v-2h26zm0-4h-2v-2h2zm0-4h-2v-2h2zm0-4h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V7h2zm0-4h-2V3h2zm-4 0h-2V3h2zm-4 0h-2V3h2zm-4 0h-2V3h2zM9 5H7V3h2zm4 0h-2V3h2zM5 25H3v-2h2zm0-4H3v-2h2zm0-4H3v-2h2zm0-4H3v-2h2zm0-4H3V7h2zm0-4H3V3h2zm3 5h10v2H8zm0 5h6v2H8z"/>`),
		g.Group(children),
	)
}