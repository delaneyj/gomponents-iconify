package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulletsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56 128a16 16 0 1 1-16-16a16 16 0 0 1 16 16ZM40 48a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm0 128a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm176-64H88a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h128a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm0-64H88a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h128a8 8 0 0 0 8-8V56a8 8 0 0 0-8-8Zm0 128H88a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h128a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}