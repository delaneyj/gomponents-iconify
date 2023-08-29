package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YellowHexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="M35.024 8.261a1.953 1.953 0 0 1 1.952 0l22.558 13.025c.605.348.976.992.976 1.69v26.048a1.95 1.95 0 0 1-.976 1.69L36.976 63.739a1.953 1.953 0 0 1-1.952 0L12.465 50.714a1.95 1.95 0 0 1-.975-1.69V22.976c0-.697.371-1.342.976-1.69L35.023 8.261Z"/><path fill="none" stroke="#000" stroke-width="2" d="M35.024 8.261a1.953 1.953 0 0 1 1.952 0l22.558 13.025c.605.348.976.992.976 1.69v26.048a1.95 1.95 0 0 1-.976 1.69L36.976 63.739a1.953 1.953 0 0 1-1.952 0L12.465 50.714a1.95 1.95 0 0 1-.975-1.69V22.976c0-.697.371-1.342.976-1.69L35.023 8.261Z"/>`),
		g.Group(children),
	)
}