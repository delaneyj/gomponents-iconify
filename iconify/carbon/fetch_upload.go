package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FetchUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 7L6 17l1.41 1.41L15 10.83V28H2v2h13a2 2 0 0 0 2-2V10.83l7.59 7.58L26 17Z"/><path fill="currentColor" d="M6 8V4h20v4h2V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4Z"/>`),
		g.Group(children),
	)
}