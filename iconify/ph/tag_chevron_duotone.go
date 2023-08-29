package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagChevronDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m230.66 132.44l-42.67 64a8 8 0 0 1-6.66 3.56H24l46.65-67.56a8 8 0 0 0 0-8.88L24 56h157.33a8 8 0 0 1 6.67 3.56l42.67 64a8 8 0 0 1-.01 8.88Z" opacity=".2"/><path d="m237.3 119.12l-42.66-64A16 16 0 0 0 181.33 48H24a8 8 0 0 0-6.58 12.55L64 128l-46.58 67.45A8 8 0 0 0 24 208h157.33a16 16 0 0 0 13.31-7.12l42.66-64a16 16 0 0 0 0-17.76Zm-56 72.88H39.25l38-55l.08-.1a16 16 0 0 0 0-17.76l-.08-.1l-38-55h142.08L224 128Z"/></g>`),
		g.Group(children),
	)
}