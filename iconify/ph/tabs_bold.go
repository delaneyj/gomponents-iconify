package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M255.49 164.55v-.06a.29.29 0 0 0 0-.09L233.2 90.25A19.87 19.87 0 0 0 214.05 76H204a12 12 0 0 0 0 24h7.07l16.8 56h-30.94L177.2 90.25A19.87 19.87 0 0 0 158.05 76H148a12 12 0 0 0 0 24h7.07l16.8 56h-30.94L121.2 90.25A19.87 19.87 0 0 0 102.05 76H42a19.87 19.87 0 0 0-19.2 14.25L.55 164.4a.29.29 0 0 0 0 .09v.06a5.68 5.68 0 0 0-.16.62A12 12 0 0 0 12 180h232a12 12 0 0 0 11.49-15.45ZM44.93 100h54.14l16.8 56H28.13Z"/>`),
		g.Group(children),
	)
}