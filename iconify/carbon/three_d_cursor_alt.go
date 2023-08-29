package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDCursorAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4z" fill="currentColor"/><path d="M15 7h2v7h-2z" fill="currentColor"/><path d="M7 15h7v2H7z" fill="currentColor"/><path d="M15 18h2v7h-2z" fill="currentColor"/><path d="M18 15h7v2h-7z" fill="currentColor"/>`),
		g.Group(children),
	)
}