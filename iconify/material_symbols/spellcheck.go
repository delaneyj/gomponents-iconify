package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spellcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.1 22l-4.25-4.25l1.4-1.4l2.85 2.85l5.65-5.65l1.4 1.4L14.1 22ZM3 16L7.85 3h2.35l4.85 13h-2.3l-1.15-3.3H6.35L5.2 16H3Zm4.05-5.2h3.9l-1.9-5.4h-.1l-1.9 5.4Z"/>`),
		g.Group(children),
	)
}