package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTeardropBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 20A104.11 104.11 0 0 0 28 124v84.33A19.69 19.69 0 0 0 47.67 228H132a104 104 0 0 0 0-208Zm0 184H52v-80a80 80 0 1 1 80 80Z"/>`),
		g.Group(children),
	)
}