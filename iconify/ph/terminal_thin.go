package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116 128a4 4 0 0 1-1.34 3l-72 64a4 4 0 1 1-5.32-6L106 128L37.34 67a4 4 0 0 1 5.32-6l72 64a4 4 0 0 1 1.34 3Zm100 60h-96a4 4 0 0 0 0 8h96a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}