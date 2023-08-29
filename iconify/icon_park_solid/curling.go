package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCurling0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M38 42H10c-3.31 0-6-2.68-6-6v-8c0-3.31 2.68-6 6-6h28c3.31 0 6 2.68 6 6v8c0 3.32-2.68 6-6 6Z"/><path stroke-linecap="round" d="M4 32h40"/><path fill="#fff" stroke-linecap="round" d="m10 22l6.48-14.8A1.98 1.98 0 0 1 18.31 6H35.5C37.43 6 39 7.57 39 9.5S37.43 13 35.5 13H22v9H10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCurling0)"/>`),
		g.Group(children),
	)
}