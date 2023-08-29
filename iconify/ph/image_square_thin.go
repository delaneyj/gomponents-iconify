package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a12 12 0 0 0-12 12v160a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12ZM44 208V48a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4v87l-31.51-31.52a12 12 0 0 0-17 0L55 212h-7a4 4 0 0 1-4-4Zm164 4H66.34l102.83-102.83a4 4 0 0 1 5.66 0L212 146.34V208a4 4 0 0 1-4 4ZM96 116a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm0-32a12 12 0 1 1-12 12a12 12 0 0 1 12-12Z"/>`),
		g.Group(children),
	)
}