package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 72C154.542 72 72 154.542 72 256s82.542 184 184 184s184-82.542 184-184S357.458 72 256 72Zm0 336c-83.813 0-152-68.187-152-152s68.187-152 152-152s152 68.187 152 152s-68.187 152-152 152Z"/>`),
		g.Group(children),
	)
}