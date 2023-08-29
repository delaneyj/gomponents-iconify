package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM3 7.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}