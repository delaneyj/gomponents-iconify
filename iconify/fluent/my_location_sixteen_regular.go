package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyLocationSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm.5-9a.5.5 0 0 0-1 0v1.525A5.002 5.002 0 0 0 3.025 7.5H1.5a.5.5 0 0 0 0 1h1.525A5.002 5.002 0 0 0 7.5 12.976V14.5a.5.5 0 0 0 1 0v-1.524A5.002 5.002 0 0 0 12.975 8.5H14.5a.5.5 0 1 0 0-1h-1.525A5.002 5.002 0 0 0 8.5 3.025V1.5ZM8 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}