package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePencil0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePencil1" fill="currentColor"><path id="fePencil2" d="M3 18L15 6l3 3L6 21H3v-3ZM16 5l2-2l3 3l-2.001 2.001L16 5Z"/></g></g>`),
		g.Group(children),
	)
}