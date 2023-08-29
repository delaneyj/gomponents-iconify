package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePicture0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePicture1" fill="currentColor" fill-rule="nonzero"><path id="fePicture2" d="M4 6v12h16V6H4Zm0-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm3.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM7 14l2-2l2 2l4-4l3 3v3H7v-2Z"/></g></g>`),
		g.Group(children),
	)
}