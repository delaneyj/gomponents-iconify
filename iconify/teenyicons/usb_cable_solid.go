package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbCableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v6h7V0ZM6 4V2h1v2H6Zm2 0V2h1v2H8Z" clip-rule="evenodd"/><path fill="currentColor" d="M12 7H3v3.5A1.5 1.5 0 0 0 4.5 12H5v2h1v1h1v-1h1v1h1v-1h1v-2h.5a1.5 1.5 0 0 0 1.5-1.5V7Z"/>`),
		g.Group(children),
	)
}