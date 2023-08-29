package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsNineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M70 60a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm58-10a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm68 20a10 10 0 1 0-10-10a10 10 0 0 0 10 10ZM60 118a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm68 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm68 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10ZM60 186a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm68 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm68 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Z"/>`),
		g.Group(children),
	)
}