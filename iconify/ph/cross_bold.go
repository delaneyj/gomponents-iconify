package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 68h-36V32a20 20 0 0 0-20-20h-32a20 20 0 0 0-20 20v36H56a20 20 0 0 0-20 20v32a20 20 0 0 0 20 20h36v84a20 20 0 0 0 20 20h32a20 20 0 0 0 20-20v-84h36a20 20 0 0 0 20-20V88a20 20 0 0 0-20-20Zm-4 48h-36a20 20 0 0 0-20 20v84h-24v-84a20 20 0 0 0-20-20H60V92h36a20 20 0 0 0 20-20V36h24v36a20 20 0 0 0 20 20h36Z"/>`),
		g.Group(children),
	)
}