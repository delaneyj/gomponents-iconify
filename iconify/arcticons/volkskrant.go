package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volkskrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.431 5.5l-15.417 37l-15.417-37m19.594 26.976L16.951 5.5m18.882 0H42m-37 0h15.548"/>`),
		g.Group(children),
	)
}