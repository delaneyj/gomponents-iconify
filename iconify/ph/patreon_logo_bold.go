package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M170 36a73.68 73.68 0 0 0-50 19.53A20 20 0 0 0 100 36H64a20 20 0 0 0-20 20v152a20 20 0 0 0 20 20h36a20 20 0 0 0 20-20v-43.5A74 74 0 1 0 170 36ZM96 204H68V60h28Zm74-44a50 50 0 1 1 50-50a50.06 50.06 0 0 1-50 50Z"/>`),
		g.Group(children),
	)
}