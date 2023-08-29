package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxListDetail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10M11 9h5"/><rect width="4" height="4" x="3" y="5" fill="currentColor" rx="1"/><path d="M11 15h10m-10 4h5"/><rect width="4" height="4" x="3" y="15" fill="currentColor" rx="1"/></g>`),
		g.Group(children),
	)
}