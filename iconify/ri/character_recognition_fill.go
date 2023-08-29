package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterRecognitionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.998 3v18h-18V3h18Zm-8.001 3h-2L6.598 17h2.154l1.199-3h4.09l1.201 3h2.155l-4.4-11Zm-1 2.885L13.242 12H10.75l1.247-3.115Z"/>`),
		g.Group(children),
	)
}