package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevModePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m12 19.01l.01-.011"/><path d="M18 18v3.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V18M18 6V2.6a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6V6"/><path stroke-linejoin="round" d="M15.5 8.5L19 12l-3.5 3.5m-7-7L5 12l3.5 3.5"/></g>`),
		g.Group(children),
	)
}