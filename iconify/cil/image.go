package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 472h432V40H40Zm400-123.858L328.628 236.769l46.6-46.6L440 254.935ZM72 72h368v137.68l-64.769-64.77L306 214.142l-100-100l-134 134Zm0 221.4l134-134l234 234V440H72Z"/>`),
		g.Group(children),
	)
}