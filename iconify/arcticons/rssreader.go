package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rssreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 13.573A28.927 28.927 0 0 1 34.427 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5a37 37 0 0 1 37 37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 18.282A24.218 24.218 0 0 1 29.718 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.318 42.5A16.818 16.818 0 0 0 5.5 25.682v-7.4"/><path fill="none" stroke="currentColor" d="M5.5 13.573V5.5m16.818 37h7.4m4.709 0H42.5"/><circle cx="10.246" cy="37.755" r="4.745" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}