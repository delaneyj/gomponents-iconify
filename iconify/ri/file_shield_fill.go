package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileShieldFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7v13.992A1 1 0 0 1 20.007 22H3.993A.993.993 0 0 1 3 21.008V2.992C3 2.455 3.447 2 3.998 2h11.999L21 7ZM8 8v5.6c0 .85.446 1.643 1.187 2.114L12 17.5l2.813-1.786A2.51 2.51 0 0 0 16 13.6V8H8Zm2 2h4v3.6c0 .158-.09.318-.26.425L12 15.131l-1.74-1.106A.512.512 0 0 1 10 13.6V10Z"/>`),
		g.Group(children),
	)
}