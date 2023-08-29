package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneCameraFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.803 4A6 6 0 0 0 23 12.197V19a.996.996 0 0 1-1.003 1.001H2.002A1 1 0 0 1 1 19V5c0-.552.44-1 1.002-1h12.8ZM20 11a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-1 6v3h2v-3h-2Z"/>`),
		g.Group(children),
	)
}