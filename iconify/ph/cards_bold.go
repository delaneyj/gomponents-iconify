package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 72H36a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V92a20 20 0 0 0-20-20Zm-4 128H40V96h136Zm64-148v124a12 12 0 0 1-24 0V56H64a12 12 0 0 1 0-24h156a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}