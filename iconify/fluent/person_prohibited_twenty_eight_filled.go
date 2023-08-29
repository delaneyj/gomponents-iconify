package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonProhibitedTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.114 25.719A7.475 7.475 0 0 1 13 20.5c0-1.688.558-3.247 1.5-4.5H5a3 3 0 0 0-3 3v.715C2 23.433 6.21 26 12 26a16.81 16.81 0 0 0 3.114-.281ZM18 8A6 6 0 1 0 6 8a6 6 0 0 0 12 0Zm9 12.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Zm-1.5 0c0-1.11-.362-2.136-.974-2.965l-6.991 6.991A5 5 0 0 0 25.5 20.5Zm-9.026 2.965l6.991-6.991a5 5 0 0 0-6.991 6.991Z"/>`),
		g.Group(children),
	)
}