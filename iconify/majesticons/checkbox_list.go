package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10m-10 7h10m-10 7h10"/><rect width="4" height="4" x="3" y="3" fill="currentColor" rx="1"/><rect width="4" height="4" x="3" y="10" fill="currentColor" rx="1"/><rect width="4" height="4" x="3" y="17" fill="currentColor" rx="1"/></g>`),
		g.Group(children),
	)
}