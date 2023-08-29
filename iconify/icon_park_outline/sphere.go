package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 32c11.046 0 20-3.582 20-8s-8.954-8-20-8s-20 3.582-20 8s8.954 8 20 8Z"/><path stroke-linecap="round" d="M32 24c0 11.046-3.582 20-8 20s-8-8.954-8-20s3.582-20 8-20s8 8.954 8 20Z"/><circle cx="24" cy="24" r="20"/></g>`),
		g.Group(children),
	)
}