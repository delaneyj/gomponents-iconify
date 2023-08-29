package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carrot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 21s9.834-3.489 12.684-6.34a4.487 4.487 0 0 0 0-6.344a4.483 4.483 0 0 0-6.342 0c-2.86 2.861-6.347 12.689-6.347 12.689zm6-8l-1.5-1.5M16 14l-2-2m8-4s-1.14-2-3-2c-1.406 0-3 2-3 2s1.14 2 3 2s3-2 3-2z"/><path d="M16 2s-2 1.14-2 3s2 3 2 3s2-1.577 2-3c0-1.86-2-3-2-3z"/></g>`),
		g.Group(children),
	)
}