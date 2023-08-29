package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 19a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0V5.5a25.49 25.49 0 0 0 12.817-3.457l.933-.543h.25v3.85m0 11.65a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0V5.35M8.5 9.5a25.49 25.49 0 0 0 12.817-3.457l.933-.543l.25-.15"/>`),
		g.Group(children),
	)
}