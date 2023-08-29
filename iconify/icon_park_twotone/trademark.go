package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trademark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTrademark0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#555"/><path d="M12 19h8m-4 0v10m10 0V19l5 6l5-6v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTrademark0)"/>`),
		g.Group(children),
	)
}