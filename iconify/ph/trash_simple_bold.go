package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 48H40a12 12 0 0 0 0 24h4v136a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20V72h4a12 12 0 0 0 0-24Zm-28 156H68V72h120ZM76 20A12 12 0 0 1 88 8h80a12 12 0 0 1 0 24H88a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}