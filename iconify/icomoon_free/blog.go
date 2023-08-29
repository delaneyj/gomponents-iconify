package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 0v1.5a8.46 8.46 0 0 1 6.01 2.489a8.472 8.472 0 0 1 2.489 6.01h1.5c0-5.523-4.477-10-10-10z"/><path fill="currentColor" d="M6 3v1.5c1.469 0 2.85.572 3.889 1.611S11.5 8.531 11.5 10H13a7 7 0 0 0-7-7zm1.5 3l-1 1L3 8l-3 6.5l.396.396l3.638-3.638a1 1 0 1 1 .707.707l-3.638 3.638l.396.396l6.5-3l1-3.5l1-1l-2.5-2.5z"/>`),
		g.Group(children),
	)
}