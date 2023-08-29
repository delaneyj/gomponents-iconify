package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 3v26h-2V3zm-4 0v2h-2V3zm-4 0v2h-2V3zm-4 0v2h-2V3zm-4 0v2h-2V3zM9 3v2H7V3zM5 3v2H3V3zm0 4v2H3V7zm0 4v2H3v-2zm0 4v2H3v-2zm0 8v2H3v-2zm0-4v2H3v-2zm20 8v2h-2v-2zm-4 0v2h-2v-2zm-4 0v2h-2v-2zm-4 0v2h-2v-2zm-4 0v2H7v-2zm-4 0v2H3v-2zm3-17h10v2H8zm0 5h6v2H8z"/>`),
		g.Group(children),
	)
}