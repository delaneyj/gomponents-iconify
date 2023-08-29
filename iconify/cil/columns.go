package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 48H40a24.028 24.028 0 0 0-24 24v376a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24Zm-8 32v64H48V80ZM48 176h192v264H48Zm224 264V176h192v264Z"/>`),
		g.Group(children),
	)
}