package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HacEnc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.681 5.5H41.84v37H30.681zM17.319 16.549V5.5H6.161v37h11.158V31.531"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.16 42.5c4.895-8.76 13.068-13.9 24.52-13.9M6.16 32.956c5.237-7.219 13.068-11.305 24.52-11.305M6.16 23.413c5.58-5.677 13.068-8.712 24.52-8.712"/>`),
		g.Group(children),
	)
}