package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileShieldTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 9V4H5v16h6.056a4.99 4.99 0 0 0 1.18 1.085l1.39.915H3.993A.993.993 0 0 1 3 21.008V2.992C3 2.455 3.449 2 4.002 2h10.995L21 8v1h-7Zm-2 2h9v5.949c0 .99-.501 1.916-1.336 2.465L16.5 21.498l-3.164-2.084A2.954 2.954 0 0 1 12 16.95V11Zm2 5.949c0 .316.162.614.436.795l2.064 1.36l2.064-1.36a.954.954 0 0 0 .436-.795V13h-5v3.949Z"/>`),
		g.Group(children),
	)
}