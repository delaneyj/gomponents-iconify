package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9v3.5A3.5 3.5 0 0 0 5.5 16h1a3.5 3.5 0 0 0 3.5-3.5V9H2zm6-2V1a1 1 0 1 1 2 0v6a2 2 0 0 1 2 2v3.5a5.502 5.502 0 0 1-4.155 5.334A1.5 1.5 0 0 1 6.5 20h-1a1.5 1.5 0 0 1-1.345-2.166A5.502 5.502 0 0 1 0 12.5V9a2 2 0 0 1 2-2V1a1 1 0 1 1 2 0v6h4z"/>`),
		g.Group(children),
	)
}