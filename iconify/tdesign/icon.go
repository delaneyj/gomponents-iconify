package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H8v5.15A7.5 7.5 0 1 0 16.85 16H22V2Zm-5.016 12A7.501 7.501 0 0 0 10 7.016V4h10v10h-3.016ZM9.5 9a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Z"/>`),
		g.Group(children),
	)
}