package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurntableLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12c0-4.714 0-7.071 1.172-8.536C4.343 2 6.229 2 10 2h4c3.771 0 5.657 0 6.828 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.172 8.535C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.465C2 19.072 2 16.714 2 12Z"/><path stroke-linecap="round" d="m15 17l.894-.447A2 2 0 0 0 17 14.763V10.5"/><rect width="8" height="8" x="6" y="8" rx="4"/><path d="M15.5 9a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/></g>`),
		g.Group(children),
	)
}