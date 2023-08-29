package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strava(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.731 0L2 9.125h2.788L6.73 5.497l1.93 3.628h2.766L6.731 0zm4.694 9.125l-1.372 2.756L8.66 9.125H6.547L10.053 16l3.484-6.875h-2.112z"/>`),
		g.Group(children),
	)
}