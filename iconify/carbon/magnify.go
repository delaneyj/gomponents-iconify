package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs/><path d="M12 21H4V4h18v8h2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v17a2 2 0 0 0 2 2h8z" fill="currentColor"/><path d="M30 28.58l-3.11-3.11a6 6 0 1 0-1.42 1.42L28.58 30zM22 26a4 4 0 1 1 4-4a4 4 0 0 1-4 4z" fill="currentColor"/>`),
		g.Group(children),
	)
}