package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm8.5 0a.5.5 0 0 0-.254-.435L7.62 6.077a.75.75 0 0 0-1.12.652v2.542a.75.75 0 0 0 1.12.653l2.626-1.489A.5.5 0 0 0 10.5 8Z"/>`),
		g.Group(children),
	)
}