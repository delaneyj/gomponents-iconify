package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAreaMap0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m18 24l-6 9l-8-4.5V42h40V15l-7 8l-6-5l-7 8l-6-2Z"/><path d="M4 28.5V17l7 6l5.5-8l6 3L31 9l5.5 4.5L44 6v9.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAreaMap0)"/>`),
		g.Group(children),
	)
}