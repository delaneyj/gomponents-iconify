package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsNineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 48v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M72 60a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm56-12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 24a12 12 0 1 0-12-12a12 12 0 0 0 12 12ZM60 116a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12ZM60 184a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/></g>`),
		g.Group(children),
	)
}