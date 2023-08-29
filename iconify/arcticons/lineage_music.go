package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.881" cy="31.876" r="7.381" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.119" cy="31.876" r="7.381" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.262 31.876v-14.54a1.897 1.897 0 0 1 1.898-1.897h18.443a1.897 1.897 0 0 1 1.897 1.898v14.539"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.262 18.445v-5.907a3.795 3.795 0 0 1 3.795-3.795h14.648a3.795 3.795 0 0 1 3.795 3.795v5.907"/>`),
		g.Group(children),
	)
}