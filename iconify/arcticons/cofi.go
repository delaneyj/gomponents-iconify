package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cofi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.714 23.744c.555 5.063-1.294 9.488-4.511 11.887a3.405 3.405 0 0 0-4.189 1.778a9.964 9.964 0 0 1-1.592.126c-6.518 0-12.48-6.174-13.315-13.79c-.567-5.173 1.376-9.68 4.722-12.04c.896.754 3.69-.234 3.723-1.58a9.92 9.92 0 0 1 1.847-.172c6.519 0 12.48 6.175 13.315 13.791Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.894 12.105c2.752 5.058 5.61 19.832 9.539 24.19M20.154 11.14c2.782 5.057 5.67 19.828 9.64 24.184M24 2.5v4.3m0 34.4v4.3M15.772 4.137l1.646 3.972M30.582 39.89l1.646 3.973M8.797 8.797l3.04 3.04m24.325 24.325l3.04 3.04M4.137 15.772l3.972 1.646M39.89 30.582l3.973 1.646M2.5 24h4.3m34.4 0h4.3M4.137 32.228l3.972-1.646M39.89 17.418l3.973-1.646M8.797 39.203l3.04-3.04m24.325-24.325l3.04-3.04m-23.43 35.065l1.646-3.972M30.582 8.11l1.646-3.973"/>`),
		g.Group(children),
	)
}