package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTShirt0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m9 9l9-5h12l9 5l4 15l-8 6v14H13V30l-8-6L9 9Z"/><path d="M13 31v-7m22 7v-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTShirt0)"/>`),
		g.Group(children),
	)
}