package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2042 1600l-317 317l-90-90l163-163h-518v-128h518l-163-163l90-90l317 317zm-666 192l127 128H677l-549-549V128h1792v1123l-128-128V256H256v1024h512v512h608zm-736-384H347l293 293v-293z"/>`),
		g.Group(children),
	)
}