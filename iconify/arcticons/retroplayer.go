package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retroplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.255 20.667L12.62 5.034a3.85 3.85 0 0 0-5.8 3.333v31.266a3.85 3.85 0 0 0 5.8 3.333l26.635-15.633a3.849 3.849 0 0 0 0-6.667Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.367 31.106h10.789c3.791 0 6.865-3.637 6.865-8.123s-3.074-8.122-6.865-8.122H7.366"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.822 14.925h4.28l11.632 22.107h0"/>`),
		g.Group(children),
	)
}