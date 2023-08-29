package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenAnAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOpenAnAccount0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 24V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h15m15.05-9v10M44 37.95H34"/><circle cx="24" cy="18" r="5" fill="#555"/><path d="M33 31c0-4.418-4.03-8-9-8s-9 3.582-9 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOpenAnAccount0)"/>`),
		g.Group(children),
	)
}