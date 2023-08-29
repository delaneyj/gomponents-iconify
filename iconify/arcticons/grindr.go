package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grindr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.235 26.628c0 1.302 1.695 4.543 4.543 4.543s6.263-2.554 6.263-5.452c0-2.161-2.014-2.48-3.045-2.48c-.492 0-7.761.294-7.761 3.389Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c-6.09 0-13.209-5.896-13.532-6.46C6.12 29.478 5.63 7.35 5.63 7.35L8.582 4.5A112.916 112.916 0 0 0 24 5.974A112.916 112.916 0 0 0 39.418 4.5l2.952 2.849s-.491 22.128-4.838 29.692C37.209 37.604 30.09 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.765 26.628c0 1.302-1.695 4.543-4.543 4.543s-6.263-2.554-6.263-5.452c0-2.161 2.014-2.48 3.045-2.48c.492 0 7.761.294 7.761 3.389Z"/>`),
		g.Group(children),
	)
}