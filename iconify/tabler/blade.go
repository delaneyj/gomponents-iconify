package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17.707 3.707l2.586 2.586a1 1 0 0 1 0 1.414l-.586.586a1 1 0 0 0 0 1.414l.586.586a1 1 0 0 1 0 1.414l-8.586 8.586a1 1 0 0 1-1.414 0l-.586-.586a1 1 0 0 0-1.414 0l-.586.586a1 1 0 0 1-1.414 0l-2.586-2.586a1 1 0 0 1 0-1.414l.586-.586a1 1 0 0 0 0-1.414l-.586-.586a1 1 0 0 1 0-1.414l8.586-8.586a1 1 0 0 1 1.414 0l.586.586a1 1 0 0 0 1.414 0l.586-.586a1 1 0 0 1 1.414 0zM8 16l3.2-3.2m1.6-1.6L16 8m-2 0l2 2m-8 4l2 2"/><path d="M11 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}