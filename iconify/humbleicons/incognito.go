package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Incognito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M10 15.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm9 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 15l.211-.106a4 4 0 0 1 3.578 0L14 15m3-6l-1.65-4.125a1 1 0 0 0-1.245-.577l-1.473.491a2 2 0 0 1-1.264 0L9.894 4.3a1 1 0 0 0-1.245.576L7 9m-4 1c7.5-1 11-1 18 0"/></g>`),
		g.Group(children),
	)
}