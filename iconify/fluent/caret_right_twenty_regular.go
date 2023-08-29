package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.628 14.982A1 1 0 0 1 7 14.204V5.797a1 1 0 0 1 1.628-.778l4.723 3.814a1.5 1.5 0 0 1 0 2.334l-4.723 3.815ZM8 5.797v8.407l4.723-3.815a.5.5 0 0 0 0-.778L8 5.797Z"/>`),
		g.Group(children),
	)
}