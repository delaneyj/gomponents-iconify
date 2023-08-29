package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuarantinePlaceTimeCalendarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 18.373a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428Zm-.452-7.464h.904m-.452 0v2.036m3.039-.965l.64.64m-.32-.32l-1.44 1.44m2.831 1.467v.904m0-.452h-2.036m.965 3.039l-.64.64m.32-.32l-1.44-1.44m-1.467 2.831h-.904m.452 0v-2.036m-3.039.965l-.64-.64m.32.32l1.44-1.44M7.25 16.111v-.904m0 .452h2.036m-.965-3.039l.64-.64m-.32.32l1.44 1.44"/><path stroke-linecap="round" d="M21.75 2.82H2.25a1.5 1.5 0 0 0-1.5 1.5v17.375a1.5 1.5 0 0 0 1.5 1.5h19.5a1.5 1.5 0 0 0 1.5-1.5V4.32a1.5 1.5 0 0 0-1.5-1.5Z"/><path d="M.75 8.32h22.5"/><path stroke-linecap="round" d="M7.25 4.805v-4m9 4v-4"/></g>`),
		g.Group(children),
	)
}