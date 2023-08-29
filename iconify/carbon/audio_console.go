package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioConsole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 14.005h-1v-12h-2v12h-1a2.002 2.002 0 0 0-2 2v2a2.002 2.002 0 0 0 2 2h1v10h2v-10h1a2.003 2.003 0 0 0 2-2v-2a2.002 2.002 0 0 0-2-2zm0 4h-4v-2h4zm-10-12h-1v-4h-2v4h-1a2.002 2.002 0 0 0-2 2v2a2.002 2.002 0 0 0 2 2h1v18h2v-18h1a2.002 2.002 0 0 0 2-2v-2a2.002 2.002 0 0 0-2-2zm0 4h-4v-2h4zm-10 10H7v-18H5v18H4a2.002 2.002 0 0 0-2 2v2a2.002 2.002 0 0 0 2 2h1v4h2v-4h1a2.002 2.002 0 0 0 2-2v-2a2.002 2.002 0 0 0-2-2zm0 4H4v-2h4z"/>`),
		g.Group(children),
	)
}