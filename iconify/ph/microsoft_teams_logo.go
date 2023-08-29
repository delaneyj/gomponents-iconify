package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftTeamsLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 104a8 8 0 0 1-8 8h-8v40a8 8 0 0 1-16 0v-40h-8a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Zm120-11.26V152a40 40 0 0 1-36.63 39.85a64 64 0 0 1-118.7.15H40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h56.81a40 40 0 0 1 73.31-28.85A32 32 0 0 1 211.69 80h7.57A12.76 12.76 0 0 1 232 92.74ZM112 56a23.82 23.82 0 0 0 1.38 8H136a16 16 0 0 1 15.07 10.68A24 24 0 1 0 112 56Zm24 120V80H40v96h96Zm48-80h-32v80a16 16 0 0 1-16 16H94.44A48 48 0 0 0 184 168Zm16-32a16 16 0 0 0-24.4-13.6A39.89 39.89 0 0 1 168 80h16a16 16 0 0 0 16-16Zm16 32h-16v72a62.76 62.76 0 0 1-.36 6.75A24 24 0 0 0 216 152Z"/>`),
		g.Group(children),
	)
}