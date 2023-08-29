package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binocular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.5 14L20 9s-.5-2-2.5-2c0 0 0-2-2-2s-2 2-2 2h-3s0-2-2-2s-2 2-2 2C4.5 7 4 9 4 9l-1.5 5"/><path d="M6 20a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm12 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M12 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`),
		g.Group(children),
	)
}