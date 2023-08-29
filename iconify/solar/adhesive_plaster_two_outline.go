package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdhesivePlasterTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.879 11.293a1 1 0 1 1-1.414 1.414a1 1 0 0 1 1.414-1.414Zm2.828 4.243a1 1 0 1 0-1.414-1.415a1 1 0 0 0 1.414 1.415Zm0-7.071a1 1 0 1 1-1.414 1.414a1 1 0 0 1 1.414-1.414Zm2.828 4.242a1 1 0 1 0-1.414-1.414a1 1 0 0 0 1.414 1.414Z"/><path fill-rule="evenodd" d="M3.054 3.054a6.16 6.16 0 0 1 8.711 0l9.18 9.18a6.16 6.16 0 0 1-8.71 8.712l-9.18-9.18a6.16 6.16 0 0 1 0-8.712Zm7.078.573a4.66 4.66 0 0 0-6.505 6.505l6.505-6.505Zm1.103 1.018l-6.59 6.59l8.12 8.12l6.59-6.59l-8.12-8.12Zm9.138 9.223l-6.505 6.505a4.66 4.66 0 0 0 6.505-6.505Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}