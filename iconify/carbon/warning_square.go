package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 20a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 16 20zM15 9h2v9h-2z"/><path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v20a2.002 2.002 0 0 1-2 2ZM6 6v20h20.001L26 6Z"/>`),
		g.Group(children),
	)
}