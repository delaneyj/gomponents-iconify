package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TramThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 52h-52V20h36a4 4 0 0 0 0-8H88a4 4 0 0 0 0 8h36v32H72a28 28 0 0 0-28 28v104a28 28 0 0 0 28 28h16l-19.2 25.6a4 4 0 1 0 6.4 4.8L98 212h60l22.8 30.4a4 4 0 0 0 6.4-4.8L168 212h16a28 28 0 0 0 28-28V80a28 28 0 0 0-28-28ZM72 60h112a20 20 0 0 1 20 20v44H52V80a20 20 0 0 1 20-20Zm112 144H72a20 20 0 0 1-20-20v-52h152v52a20 20 0 0 1-20 20Zm-92-32a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm88 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}