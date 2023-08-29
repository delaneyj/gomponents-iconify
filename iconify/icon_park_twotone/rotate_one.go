package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRotateOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M44 24c0-11.046-8.954-20-20-20S4 12.954 4 24s8.954 20 20 20c6.957 0 13.084-3.552 16.667-8.94M44 24H30"/><circle cx="24" cy="24" r="6" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRotateOne0)"/>`),
		g.Group(children),
	)
}