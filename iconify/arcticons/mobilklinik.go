package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilklinik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.692 32.112c.977 1.727-1.085 4.065-6.692 5.43c0 0-2.007-3.884-.961-6.208s3.03-1.296 2.702 1.679c0 0 3.974-2.629 4.95-.902Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.148 33.846c3.158-8.134 6.293-16.275 9.465-24.403a7.21 7.21 0 0 1 9.548-3.56a6.711 6.711 0 0 1 3.206 8.907c-3.102 7.99-6.182 15.99-9.299 23.976a7.21 7.21 0 0 1-9.548 3.56a6.664 6.664 0 0 1-3.372-8.48Zm1.768-6.011l-4.64-9.536M19.9 11.95A6.951 6.951 0 1 1 12.95 5a6.95 6.95 0 0 1 6.95 6.95Z"/>`),
		g.Group(children),
	)
}