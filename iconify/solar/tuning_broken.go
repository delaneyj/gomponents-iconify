package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 14.5a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm-10-5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/><path stroke-linecap="round" d="M7 13v5m0 3v1m10-11V6m0-3V2m0 20v-4M7 2v4"/></g>`),
		g.Group(children),
	)
}