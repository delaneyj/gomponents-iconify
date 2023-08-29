package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePictureSquare0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePictureSquare1" fill="currentColor" fill-rule="nonzero"><path id="fePictureSquare2" d="M5 5v14h14V5H5Zm0-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Zm3.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM7 14l2-2l2 2l3-3l3 3v3H7v-3Z"/></g></g>`),
		g.Group(children),
	)
}