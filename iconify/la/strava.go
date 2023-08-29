package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strava(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.18 2L5.9 18h4.88l3.4-6.38L17.56 18h4.84L14.18 2zm8.22 16L20 22.79L17.56 18h-3.7L20 30l6.1-12h-3.7z"/>`),
		g.Group(children),
	)
}