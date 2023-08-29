package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVip0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M12 4H4l11 40h8L12 4Z"/><path stroke-linecap="round" d="M23 44L44 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVip0)"/>`),
		g.Group(children),
	)
}