package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 1C179.8 1 111.7 34.4 64.9 87.2L0 22.3V193h170.7l-60.2-60.2C145.6 91.5 197.5 65 256 65c106.1 0 192 85.9 192 192s-85.9 192-192 192c-53 0-101-21.5-135.8-56.2L75 438c46.4 46.3 110.4 75 181 75c141.4 0 256-114.6 256-256S397.4 1 256 1z"/>`),
		g.Group(children),
	)
}