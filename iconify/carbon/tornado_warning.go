package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 4h16v2H2zm2 4h12v2H4zm-2 4h12v2H2zm0 4h10v2H2zm2 4h8v2H4zm4 4h6v2H8zm4 4h4v2h-4z"/><path fill="none" d="M23.75 10h-1.5V6h1.5ZM23 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M29.912 13.935L23.628 2.371a.718.718 0 0 0-1.256 0l-6.284 11.564A.72.72 0 0 0 16.72 15h12.56a.72.72 0 0 0 .631-1.065ZM22.25 6h1.5v4h-1.5Zm.75 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}