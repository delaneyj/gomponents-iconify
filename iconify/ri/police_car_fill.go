package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoliceCarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13.5V21a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1H5v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-7.5l-1.243-.31A1 1 0 0 1 0 12.22v-.72a.5.5 0 0 1 .5-.5h1.929L4.48 6.212A2 2 0 0 1 6.319 5H8V3h3v2h2V3h3v2h1.681a2 2 0 0 1 1.838 1.212L21.572 11H23.5a.5.5 0 0 1 .5.5v.72a1 1 0 0 1-.758.97L22 13.5ZM4 15v2a1 1 0 0 0 1 1h3.245a.5.5 0 0 0 .44-.736C7.88 15.754 6.318 15 4 15Zm16 0c-2.317 0-3.879.755-4.686 2.264a.5.5 0 0 0 .441.736H19a1 1 0 0 0 1-1v-2ZM6 7l-1.451 3.629A1 1 0 0 0 5.477 12h13.046a1.001 1.001 0 0 0 .928-1.371L18 7H6Z"/>`),
		g.Group(children),
	)
}