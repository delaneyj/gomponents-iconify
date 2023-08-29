package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 56h-56v32h48v376H88V88h48V56H80a24.028 24.028 0 0 0-24 24v392a24.028 24.028 0 0 0 24 24h352a24.028 24.028 0 0 0 24-24V80a24.028 24.028 0 0 0-24-24Z"/><path fill="currentColor" d="M192 140h128a24.028 24.028 0 0 0 24-24V16H168v100a24.028 24.028 0 0 0 24 24Zm8-92h112v60H200Z"/>`),
		g.Group(children),
	)
}