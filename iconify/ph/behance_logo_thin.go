package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BehanceLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108.16 122.74A34 34 0 0 0 90 60H32a4 4 0 0 0-4 4v128a4 4 0 0 0 4 4h62a38 38 0 0 0 14.16-73.26ZM36 68h54a26 26 0 0 1 0 52H36Zm58 120H36v-60h58a30 30 0 0 1 0 60Zm70-108a4 4 0 0 1 4-4h64a4 4 0 0 1 0 8h-64a4 4 0 0 1-4-4Zm36 28a44 44 0 1 0 35.2 70.41a4 4 0 0 0-6.4-4.81a36 36 0 0 1-64.58-17.6H240a4 4 0 0 0 4-4a44.05 44.05 0 0 0-44-44Zm-35.78 40a36 36 0 0 1 71.56 0Z"/>`),
		g.Group(children),
	)
}