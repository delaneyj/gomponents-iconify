package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreenLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22" opacity=".5"/><path d="M10 22c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14"/><path d="M10 2C6.229 2 4.343 2 3.172 3.172C2 4.343 2 6.229 2 10" opacity=".5"/><path d="M14 2c3.771 0 5.657 0 6.828 1.172C22 4.343 22 6.229 22 10"/></g>`),
		g.Group(children),
	)
}