package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 4A1.5 1.5 0 0 0 6 5.5V6h3v-.5A1.5 1.5 0 0 0 7.5 4Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15ZM5 5.5v.585A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V5.5a2.5 2.5 0 0 0-5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}