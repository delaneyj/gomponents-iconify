package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volumedown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 272v150c0 9 7 15 15 15h84l169 129c27 28 47 19 47-20V147c0-38-20-47-47-19L99 257H15c-8 0-15 7-15 15zm380 237c5 3 10 5 15 5c12 0 22-6 28-17c28-47 41-97 41-150c0-52-13-103-41-151c-9-15-28-19-43-11s-20 28-12 43c23 38 32 78 32 119s-9 81-32 119c-8 15-3 35 12 43z"/>`),
		g.Group(children),
	)
}