package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 13h-5V8M5 3.5L7 2m12 1.5L17 2"/><path d="M12 22a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/></g>`),
		g.Group(children),
	)
}