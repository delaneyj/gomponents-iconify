package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodhiTimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.319 4.5c2.851 3.15 3.447 8.98 4.596 11.534s2.469.809 7.065 4.554s5.363 7.023 5.363 12.726s-8.385 7.448-9.406 7.448s-3.15.085-4.043-.936a8.372 8.372 0 0 1-4.342.938c-1.319 0-8.895-1.62-8.895-8.77s5.235-12.427 8.683-14.3s4.98-2.043 3.736-7.785"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.515 11.393c.644 6 .336 25.198.379 28.433A12.04 12.04 0 0 0 24 43.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.434 17.113c.31.666 2.856 1.88 3.455 2.124c.354.144 2.133-.905 2.689-1.576"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.502 18.906a18.888 18.888 0 0 0 6.429 5.427a37.063 37.063 0 0 0 6.185-5.096"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.328 22.04c.883 1.314 7.973 7.4 9.6 8.167c.331-.17 8.01-5.789 9.665-8.167"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.204 28.034C12.955 29.994 21.425 34.08 24 35.06c2.107-.511 10.491-6.257 12.613-8.183"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.876 33.993c2.335 1.491 9.69 4.386 13.014 4.726c3.068-.724 12.975-5.534 13.448-6.13"/>`),
		g.Group(children),
	)
}