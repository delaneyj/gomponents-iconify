package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scientist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17c-2.67 0-8 1.34-8 4v3h16v-3c0-2.66-5.33-4-8-4Z"/><circle cx="8" cy="12" r="4" fill="currentColor"/><circle cx="20" cy="5" r="1" fill="currentColor"/><circle cx="21.5" cy="3.5" r=".5" fill="currentColor"/><circle cx="20.5" cy="1.5" r=".5" fill="currentColor"/><path fill="currentColor" d="M16 7v1h1l.003 10.031C17 19.712 18 21 19.5 21s2.5-1.207 2.5-3V8h1V7Zm5 8l-.565.424a1.765 1.765 0 0 1-1.05.326H19V15h-1v-2h1v-1h-1v-2h1V9h-1V8h3Z"/>`),
		g.Group(children),
	)
}