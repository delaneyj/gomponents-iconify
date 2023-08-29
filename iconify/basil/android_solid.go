package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.016 1.646a.5.5 0 0 1 .708 0l1.105 1.106A5.621 5.621 0 0 1 11.647 2h.706a5.62 5.62 0 0 1 2.82.753l1.107-1.107a.5.5 0 1 1 .707.708l-.984.984A5.635 5.635 0 0 1 18 7.648a.353.353 0 0 1-.353.352H6.353A.353.353 0 0 1 6 7.647c0-1.728.776-3.275 2-4.31l-.984-.983a.5.5 0 0 1 0-.708ZM10.5 5.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm3.724.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z" clip-rule="evenodd"/><path fill="currentColor" d="M4 8.947a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1Zm3 0a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1H7Zm12.21 1a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0v-6ZM7.793 21a.293.293 0 0 0-.293.294v.206a1.75 1.75 0 1 0 3.5 0v-.206a.293.293 0 0 0-.293-.294H7.793Zm5.207.294c0-.163.131-.294.293-.294h2.913c.163 0 .294.131.294.294v.206a1.75 1.75 0 1 1-3.5 0v-.206Z"/>`),
		g.Group(children),
	)
}