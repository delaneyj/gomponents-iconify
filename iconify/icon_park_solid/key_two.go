package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKeyTwo0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="15" cy="33" r="8" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="m29 16l7 6m-16 4L36 8l7 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSKeyTwo0)"/>`),
		g.Group(children),
	)
}