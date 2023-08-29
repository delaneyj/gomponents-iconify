package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinwoodButterfly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.294 27.034l6.496-4.318l-3.087 6.239l4.438 3.257l4.116-14.465l-16.301 6.168l4.338 3.119z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.294 27.034v7.204l3.409-5.283"/><circle cx="10.16" cy="12.088" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 38.13h6.912a9.176 9.176 0 0 0 9.176-9.175v-5.868a9.122 9.122 0 0 0-1.761-5.389a5.682 5.682 0 1 0-5.771-3.631a9.2 9.2 0 0 0-1.644-.155H17.088a9.2 9.2 0 0 0-1.644.155a5.655 5.655 0 1 0-5.771 3.631a9.122 9.122 0 0 0-1.76 5.39v5.867a9.176 9.176 0 0 0 9.175 9.176Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.088 28.955v3.441a9.176 9.176 0 0 1-9.176 9.176H17.088a9.176 9.176 0 0 1-9.176-9.176v-3.44"/><circle cx="37.84" cy="12.088" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.84" cy="12.088" r=".75" fill="currentColor"/><circle cx="10.16" cy="12.088" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}