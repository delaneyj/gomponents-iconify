package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wggesucht(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="37.666" cy="31.887" r="4.625" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.782 42.173l3.53-6.33M5.5 15.694l27.236-9.867l9.764 7.811"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 13.844v5.344l4.111 2.878v-5.344Zm-5.86 19.238a3.828 3.828 0 0 0-3.954-3.813a3.978 3.978 0 0 0-3.53 3.954v3.53a3.797 3.797 0 0 0 3.812 3.814h0a3.797 3.797 0 0 0 3.812-3.813h-3.813M6.014 29.26L8.84 40.567l2.826-11.305l2.827 11.305l2.827-11.305"/>`),
		g.Group(children),
	)
}