package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncognitoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path stroke-linecap="round" d="M2 11h12m8 0h-4M4 11l.614-2.455c.545-2.183.818-3.274 1.632-3.91c.76-.593 1.79-.632 3.754-.635m10 7l-.614-2.455c-.546-2.183-.818-3.274-1.632-3.91c-.76-.593-1.79-.632-3.754-.635"/><path d="M10 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path stroke-linecap="round" d="m10 17.5l.658-.33a3 3 0 0 1 2.684 0l.658.33"/></g>`),
		g.Group(children),
	)
}