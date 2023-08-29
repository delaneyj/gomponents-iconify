package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PbzBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M32.803 42.5H7.5a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v25.188M8.94 15.633h30.12"/><path d="M8.94 32.367v-8.201a3.605 3.605 0 0 1 7.21 0v8.2m15.7.001v-8.201a3.605 3.605 0 0 1 7.21 0v7.29m-18.665.911v-8.201a3.605 3.605 0 0 1 7.21 0v8.2"/></g><circle cx="38.5" cy="38.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.529 38.433a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.301"/>`),
		g.Group(children),
	)
}