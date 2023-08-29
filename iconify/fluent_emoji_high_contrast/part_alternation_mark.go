package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartAlternationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.5 28.22a1.498 1.498 0 0 1-1.403-.975L20.045 8.35l-4.152 5.643c-.327.444-.85.71-1.402.712h-.005c-.55 0-1.073-.262-1.402-.704L8.861 8.334l-3.98 9.512a1.498 1.498 0 1 1-2.765-1.157L6.942 5.157a1.754 1.754 0 0 1 1.4-1.06a1.75 1.75 0 0 1 1.614.69l4.523 6.07l4.52-6.145a1.748 1.748 0 0 1 3.046.425l7.859 21.06a1.498 1.498 0 0 1-1.403 2.023Z"/>`),
		g.Group(children),
	)
}