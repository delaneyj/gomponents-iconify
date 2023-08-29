package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashCheckmarkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.874 1a1 1 0 0 0-.959.714L1.032 8.036A.75.75 0 0 0 1.75 9h1.584l-1.28 4.389c-.384 1.316 1.324 2.2 2.178 1.128l1.306-1.64a5.5 5.5 0 0 1 4.782-7.873a.797.797 0 0 0-.068-.004h-2.03l.994-2.649A1 1 0 0 0 8.28 1H3.874ZM15 10.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-2.854-1.854L9.5 11.293l-.646-.647a.5.5 0 0 0-.708.708l1 1a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708Z"/>`),
		g.Group(children),
	)
}