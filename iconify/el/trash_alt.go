package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm-93.091 224.341h186.182v57.422h217.09v105.688H289.819V281.763h217.09v-57.422zm-183.764 243.53h553.71v507.788h-553.71V467.871z"/>`),
		g.Group(children),
	)
}