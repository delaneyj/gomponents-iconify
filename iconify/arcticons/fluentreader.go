package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fluentreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.855 39.146h5.842a1.819 1.819 0 0 0 1.819-1.819V7.319A1.819 1.819 0 0 0 40.697 5.5H10.69a1.819 1.819 0 0 0-1.82 1.819v5.636"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.888 40.753h6.348a1.48 1.48 0 0 0 1.48-1.48V14.845a1.48 1.48 0 0 0-1.48-1.48H8.808a1.48 1.48 0 0 0-1.48 1.48v6.206"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.495 21.345H6.628a1.144 1.144 0 0 0-1.144 1.144v18.867A1.143 1.143 0 0 0 6.628 42.5h18.867a1.143 1.143 0 0 0 1.144-1.143h0V22.488a1.144 1.144 0 0 0-1.144-1.144Zm-12.449 4.546h6.031m-6.031 6.032h3.92m-3.92-6.032v12.064"/>`),
		g.Group(children),
	)
}