package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditName(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEditName0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#fff"/><path d="M4 41c0-8.837 8.059-16 18-16"/><path fill="#fff" d="m31 42l10-10l-4-4l-10 10v4h4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEditName0)"/>`),
		g.Group(children),
	)
}