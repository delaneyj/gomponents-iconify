package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSZeroKey0"><g fill="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke="#fff" rx="3"/><rect width="10" height="20" x="19" y="14" stroke="#000" rx="5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSZeroKey0)"/>`),
		g.Group(children),
	)
}