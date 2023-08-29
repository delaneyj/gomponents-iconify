package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.833 17.167H44.5v13.667H30.833z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.167 17.167h13.667v13.667H17.167zm-13.667 0h13.667v13.667H3.5zm37.148 8.315v.037a2.981 2.981 0 0 1-2.981 2.981h0a2.981 2.981 0 0 1-2.982-2.981V22.48a2.981 2.981 0 0 1 2.982-2.981h0a2.981 2.981 0 0 1 2.98 2.981v.037"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M11.065 24a2.25 2.25 0 0 1 0 4.5H7.352v-9h3.713a2.25 2.25 0 0 1 0 4.5Zm0 0H7.352m17.379 0a2.25 2.25 0 1 1 0 4.5H21.02v-9h3.712a2.25 2.25 0 1 1 0 4.5Zm0 0h-3.712"/>`),
		g.Group(children),
	)
}