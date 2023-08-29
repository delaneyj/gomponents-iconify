package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSApplicationTwo0"><g fill="#fff" stroke="#fff" stroke-width="4"><circle cx="34.5" cy="13.5" r="6.5"/><circle cx="34.5" cy="34.5" r="6.5"/><circle cx="13.5" cy="13.5" r="6.5"/><circle cx="13.5" cy="34.5" r="6.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSApplicationTwo0)"/>`),
		g.Group(children),
	)
}