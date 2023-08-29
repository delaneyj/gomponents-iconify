package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sagernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 6.838L36.02 35.42l-14.433-4.527l11.827-14.75l-15.731 13.57L5.5 25.407ZM28.573 33.09l-7.154 8.073l.168-10.27"/>`),
		g.Group(children),
	)
}