package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDashedDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 12h-2v-2h2v2m-4 0h-3v-2h3v2m-5 0h-3V9h2v1h1v2M12 0v2h-2V0h2m0 4v3h-2V4h2Z"/>`),
		g.Group(children),
	)
}