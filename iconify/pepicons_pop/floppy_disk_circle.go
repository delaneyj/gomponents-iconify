package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.5 6.5a2 2 0 0 1 2-2h10.586a2 2 0 0 1 1.414.586L20.914 7.5a2 2 0 0 1 .586 1.414V19.5a2 2 0 0 1-2 2h-13a2 2 0 0 1-2-2v-13Zm2 0v13h13V8.914L17.086 6.5H6.5Z"/><path d="M8 14a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h-2v-7h-6v7H8v-7Zm.5-6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V6a1 1 0 1 0-2 0v2h-4V6a1 1 0 0 0-2 0v2Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}