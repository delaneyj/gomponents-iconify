package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moviedb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.07 23.55h18.34a2.46 2.46 0 0 0 2.47-2.47V17a2.47 2.47 0 0 0-2.47-2.48h-22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.07 23.55l-8.07 6v.58a2.47 2.47 0 0 0 2.48 2.47h23.93a2.46 2.46 0 0 0 2.47-2.47V26a2.47 2.47 0 0 0-2.47-2.48Z"/><rect width="28.88" height="9.01" x="13" y="32.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.48"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.2 26.32c0-4.48 0-10.81-.07-14a53.48 53.48 0 0 1 0-5.84l9.61 7l9.58 7s-16.84 12.38-19 13.93c-.09.06-.09-.42-.12-8.09Z"/>`),
		g.Group(children),
	)
}