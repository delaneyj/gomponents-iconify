package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftTeamsLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M144 80v96a8 8 0 0 1-8 8H40a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h96a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M80 152v-40h-8a8 8 0 0 1 0-16h32a8 8 0 0 1 0 16h-8v40a8 8 0 0 1-16 0Zm152-59.26V152a40 40 0 0 1-36.63 39.85a64 64 0 0 1-118.7.15H40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h56.81a40 40 0 0 1 73.31-28.85A32 32 0 0 1 211.69 80h7.57A12.76 12.76 0 0 1 232 92.74ZM175.6 50.4A39.89 39.89 0 0 1 168 80h16a16 16 0 1 0-8.4-29.6ZM113.38 64H136a16 16 0 0 1 15.07 10.68A24 24 0 1 0 113.38 64ZM40 176h96V80H40Zm144-8V96h-32v80a16 16 0 0 1-16 16H94.44A48 48 0 0 0 184 168Zm32-72h-16v72a62.76 62.76 0 0 1-.36 6.75A24 24 0 0 0 216 152Z"/></g>`),
		g.Group(children),
	)
}