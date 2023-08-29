package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 12H72a20 20 0 0 0-20 20v45.33a20.12 20.12 0 0 0 4 12L76 116v108a20 20 0 0 0 20 20h64a20 20 0 0 0 20-20V116l20-26.67a20.12 20.12 0 0 0 4-12V32a20 20 0 0 0-20-20Zm-4 24v16H76V36Zm-20 66.67a20.12 20.12 0 0 0-4 12V220h-56V114.67a20.12 20.12 0 0 0-4-12L76 76h104ZM140 120v32a12 12 0 0 1-24 0v-32a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}