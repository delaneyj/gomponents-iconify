package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.2 21q-.925 0-1.563-.638T3 18.8V5.1q0-.725.675-1.012T4.85 4.3L7.1 6.55L5.75 7.9l.7.7L7.8 7.25l2.6 2.6l-1.35 1.35l.7.7l1.35-1.35l2.6 2.6l-1.35 1.35l.7.7l1.35-1.35l2.6 2.6l-1.35 1.35l.7.7l1.35-1.35l2 2q.5.5.213 1.175T18.9 21H5.2Zm.8-3h8.3L6 9.7V18Z"/>`),
		g.Group(children),
	)
}