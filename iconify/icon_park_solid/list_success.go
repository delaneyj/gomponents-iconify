package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListSuccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSListSuccess0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 10h24M20 24h24M20 38h24"/><circle cx="8" cy="24" r="4" fill="#fff"/><circle cx="8" cy="38" r="4" fill="#fff"/><path d="m4 10l3 3l6-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSListSuccess0)"/>`),
		g.Group(children),
	)
}