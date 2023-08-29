package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 16.09a1 1 0 1 1-2 0V10C0 4.477 4.477 0 10 0s10 4.477 10 10v3.086a1 1 0 1 1-2 0V10a8 8 0 1 0-16 0v6.09zM2 10h2a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2zm0 2v5h2v-5H2zm14-2h2a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2zm0 2v5h2v-5h-2z"/>`),
		g.Group(children),
	)
}