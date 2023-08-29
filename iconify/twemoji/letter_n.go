package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M8.591 9.156c0-1.55.992-2.418 2.326-2.418c.589 0 1.55.465 1.954 1.023L22.7 20.877h.062V9.156c0-1.55.992-2.418 2.324-2.418c1.334 0 2.326.868 2.326 2.418v17.611c0 1.551-.992 2.418-2.326 2.418c-.588 0-1.518-.465-1.953-1.022l-9.829-12.961h-.062v11.565c0 1.551-.992 2.418-2.326 2.418s-2.326-.867-2.326-2.418V9.156z"/>`),
		g.Group(children),
	)
}