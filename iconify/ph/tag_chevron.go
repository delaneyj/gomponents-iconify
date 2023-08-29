package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagChevron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m237.3 119.12l-42.66-64A16 16 0 0 0 181.33 48H24a8 8 0 0 0-6.58 12.55L64 128l-46.58 67.45A8 8 0 0 0 24 208h157.33a16 16 0 0 0 13.31-7.12l42.66-64a16 16 0 0 0 0-17.76Zm-56 72.88H39.25l38-55l.08-.1a16 16 0 0 0 0-17.76l-.08-.1l-38-55h142.08L224 128Z"/>`),
		g.Group(children),
	)
}