package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagClockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.748a.75.75 0 0 1 .75-.75h16.504a.75.75 0 0 1 .6 1.2L16.69 9.75l.939 1.252a6.502 6.502 0 0 0-6.553 5.499H4.5v4.75a.75.75 0 0 1-.648.743L3.75 22a.75.75 0 0 1-.743-.648L3 21.25V3.748ZM17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm2 5.5h-2V15a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}