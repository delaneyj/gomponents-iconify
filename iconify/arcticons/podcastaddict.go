package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podcastaddict(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.5 24a8.5 8.5 0 0 0-8.5-8.5m-2.314.319a8.501 8.501 0 1 0 10.497 10.49M15.5 43.754V24m22 0A13.5 13.5 0 0 0 24 10.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5a21.5 21.5 0 1 0 13.5 38.222V45.5h8V24A21.5 21.5 0 0 0 24 2.5Z"/>`),
		g.Group(children),
	)
}