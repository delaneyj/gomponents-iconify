package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strava(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M14.92 17.16l1.83-3.63h2.7l-4.51 8.97l-4.57-8.97h2.7l1.85 3.63m-4.29-8.5l-2.45 4.89H4.55L10.61 1.5l6.13 12.05h-3.63l-2.48-4.89z" fill="currentColor"/>`),
		g.Group(children),
	)
}