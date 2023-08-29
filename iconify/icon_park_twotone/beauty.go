package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beauty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBeauty0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 44V22c0-9.941-8.059-18-18-18S6 12.059 6 22v22"/><path fill="#555" d="M24 43c5.523 0 10-8.954 10-20H14c0 11.046 4.477 20 10 20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBeauty0)"/>`),
		g.Group(children),
	)
}