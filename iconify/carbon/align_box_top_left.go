package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBoxTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 30H6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM6 4v24h20V4Z"/><path fill="currentColor" d="M9 7h11v2H9zm0 5h7v2H9z"/>`),
		g.Group(children),
	)
}