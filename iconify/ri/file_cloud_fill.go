package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCloudFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.997 2L21 8l.001 4.26A5.466 5.466 0 0 0 17.5 11l-.221.004a5.503 5.503 0 0 0-5.127 4.205l-.016.074l-.03.02A4.75 4.75 0 0 0 10.878 22H3.993a.993.993 0 0 1-.986-.876L3 21.008V2.992c0-.498.387-.927.885-.985L4.002 2h10.995ZM17.5 13a3.5 3.5 0 0 1 3.5 3.5l-.002.103a2.75 2.75 0 0 1-.58 5.392L20.25 22h-5.5l-.168-.005a2.75 2.75 0 0 1-.579-5.393L14 16.5a3.5 3.5 0 0 1 3.5-3.5Z"/>`),
		g.Group(children),
	)
}