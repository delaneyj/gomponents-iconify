package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0C7.172 0 0 7.172 0 16s7.172 16 16 16s16-7.172 16-16S24.828 0 16 0zm9.281 18.401c-1.068 4.214-4.906 7.198-9.281 7.198s-8.214-2.984-9.281-7.198c-.104-.479.161-1.016.641-1.12c.479-.109 1.016.156 1.12.641c.906 3.411 4 5.813 7.521 5.813s6.615-2.401 7.521-5.813c.104-.484.641-.802 1.12-.641a.891.891 0 0 1 .641 1.12z"/>`),
		g.Group(children),
	)
}