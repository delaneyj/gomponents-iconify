package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 7.8L6.442 15.276c-1.456.616-2.679.925-3.668.925c-1.12 0-1.933-.392-2.437-1.177c-.317-.504-.41-1.143-.28-1.918c.13-.775.476-1.6 1.036-2.478c.467-.71 1.232-1.643 2.297-2.8a6.122 6.122 0 0 0-.784 1.848c-.28 1.195-.028 2.072.756 2.632c.373.261.886.392 1.54.392c.522 0 1.11-.084 1.764-.252L24 7.8z"/>`),
		g.Group(children),
	)
}