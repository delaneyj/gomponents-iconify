package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTopUpBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m13 11l9-9m0 0h-5.344M22 2v5.344"/><path d="M12 2C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5M22 12c0 5.523-4.477 10-10 10a9.955 9.955 0 0 1-5-1.338"/></g>`),
		g.Group(children),
	)
}