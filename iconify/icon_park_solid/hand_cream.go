package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandCream0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#fff" d="M19 17c-2.731 1.137-2.488 4.578-2 6h14c3.902-5.687-4.216-11.849-4.216-9.479c0 2.37-4.37 2.057-7.784 3.479Z"/><path d="M13 23h22v8H13z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandCream0)"/>`),
		g.Group(children),
	)
}