package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toggles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 4)"><circle cx="3.5" cy="3.5" r="3"/><path d="M6 1.5h6.5c.828 0 2 .325 2 2s-1.172 2-2 2H6m5.5 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/><path d="M9 8.5H2.5c-.828 0-2 .325-2 2s1.172 2 2 2H9"/></g>`),
		g.Group(children),
	)
}