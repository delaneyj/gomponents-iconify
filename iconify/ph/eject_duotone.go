package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 176v24a8 8 0 0 1-8 8H48a8 8 0 0 1-8-8v-24a8 8 0 0 1 8-8h160a8 8 0 0 1 8 8ZM48.23 136h159.54c7.16 0 10.89-8.27 6-13.37l-73.74-77.5a16.76 16.76 0 0 0-24.14 0l-73.74 77.5c-4.81 5.1-1.08 13.37 6.08 13.37Z" opacity=".2"/><path d="M208 160H48a16 16 0 0 0-16 16v24a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16v-24a16 16 0 0 0-16-16Zm0 40H48v-24h160ZM48.24 144h159.52a16.18 16.18 0 0 0 14.93-9.76a15.59 15.59 0 0 0-3.1-17.12l-73.73-77.51a24.76 24.76 0 0 0-35.72 0l-73.73 77.51a15.59 15.59 0 0 0-3.1 17.12A16.18 16.18 0 0 0 48.24 144Zm73.49-93.36a8.77 8.77 0 0 1 12.54 0L207.85 128H48.14Z"/></g>`),
		g.Group(children),
	)
}