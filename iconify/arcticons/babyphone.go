package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Babyphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.717 30.278a62.017 62.017 0 0 1-6.302-1.11a3.11 3.11 0 0 0-3 .58c-.53.54-2.061 2.05-3.652 3.611A28.698 28.698 0 0 1 14.66 20.245c1.55-1.59 3.001-3.11 3.591-3.64a3.111 3.111 0 0 0 .58-3.002a61.84 61.84 0 0 1-1.11-6.312a2 2 0 0 0-2.196-1.784a2.06 2.06 0 0 0-.105.014H6.897a1.5 1.5 0 0 0-1.36 1.37c-.55 7.692 3.74 15.924 4.611 17.505h0v.06l.12.23h0A35.45 35.45 0 0 0 23.272 37.69h0l.44.25h0c2 1.06 9.953 5.062 17.386 4.511a1.437 1.437 0 0 0 1.39-1.36V32.58a2.131 2.131 0 0 0-1.666-2.287a2.06 2.06 0 0 0-.105-.014Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.914 22.2a6.189 6.189 0 0 1-1.935-3.511h12.196a6.474 6.474 0 0 1-2.019 3.603M30.134 23.7a8.04 8.04 0 0 1-1.101.114a6.223 6.223 0 0 1-1.135-.114"/><circle cx="25.893" cy="24.096" r="2.001" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.094" cy="24.096" r="2.001" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.491 16.994h1l.488 1.695m6.649-7.255v5.315h5.6a6.202 6.202 0 0 0-5.6-5.315Z"/>`),
		g.Group(children),
	)
}