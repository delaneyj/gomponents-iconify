package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WritingFluently(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" d="M24 24V19L39 4L44 9L29 24H24Z" clip-rule="evenodd"/><path d="M16 24H9C6.23858 24 4 26.2386 4 29C4 31.7614 6.23858 34 9 34H39C41.7614 34 44 36.2386 44 39C44 41.7614 41.7614 44 39 44H18"/></g>`),
		g.Group(children),
	)
}