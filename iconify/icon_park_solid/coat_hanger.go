package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoatHanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCoatHanger0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 13a5 5 0 1 1 5 5v4"/><path fill="#fff" d="M44 36H4c0-5 6-14 20-14s20 9 20 14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCoatHanger0)"/>`),
		g.Group(children),
	)
}