package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WotCompendiumUnofficial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.904 10.115L5.82 41.256a37.626 37.626 0 0 1 11.237-2.006Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.392 18.182c1.675-2.77 7.454-9.03 15.44-12.682c-2.006 7.504 3.692 24.279 8.347 31.623c-14.567.441-31.743 5.377-31.743 5.377a4.647 4.647 0 0 1 0-2.504"/><circle cx="27.391" cy="26.588" r="7.163" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.391 33.752a3.582 3.582 0 0 0 0-7.164m0-7.163a3.582 3.582 0 1 0 0 7.163"/>`),
		g.Group(children),
	)
}