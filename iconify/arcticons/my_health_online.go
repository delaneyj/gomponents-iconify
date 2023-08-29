package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyHealthOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m27.62 17.316l3.616 12.641L27.033 43.5h-7.397"/><path d="M27.033 43.5h8.253l-4.05-13.543M12.714 43.5l11.525-39l3.397 13.047l-8 25.953h-6.922Z"/></g>`),
		g.Group(children),
	)
}