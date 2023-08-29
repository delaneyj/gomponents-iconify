package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlightlyFrowningFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#FFCC4D"/><path fill="#664500" d="M25.485 27.379C25.44 27.2 24.317 23 18 23c-6.318 0-7.44 4.2-7.485 4.379a.499.499 0 0 0 .237.554a.505.505 0 0 0 .6-.077C11.371 27.837 13.306 26 18 26s6.63 1.837 6.648 1.855a.502.502 0 0 0 .598.081a.5.5 0 0 0 .239-.557z"/><ellipse cx="12" cy="13.5" fill="#664500" rx="2.5" ry="3.5"/><ellipse cx="24" cy="13.5" fill="#664500" rx="2.5" ry="3.5"/>`),
		g.Group(children),
	)
}