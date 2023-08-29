package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M497 125L380 270L160 50q44-10 88-10q136 0 249 85zM356 301l74 74l-27 27l-71-71l-83 103l-1 1v-1L0 125q35-27 79-47L35 34L62 7z"/>`),
		g.Group(children),
	)
}