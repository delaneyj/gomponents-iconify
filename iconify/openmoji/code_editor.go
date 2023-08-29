package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11 16.083h50v39.833H11z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M11 16.001h50v39.998H11z"/><path d="M16.329 16.479v4.375H11h50m-32.167 9.541l-6.927 6.927m6.927 6.995l-6.927-6.928m16.278-9.237L32.809 46.25m9.35-1.998l6.927-6.928m-6.927-6.994l6.927 6.927"/></g>`),
		g.Group(children),
	)
}