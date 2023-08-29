package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosCopyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M304 96h-16v80h80v-16h-64z" fill="currentColor"/><path d="M325.3 64H160v48h-48v336h240v-48h48V139l-74.7-75zM336 432H128V128h32v272h176v32zm48-48H176V80h142.7l65.3 65.6V384z" fill="currentColor"/>`),
		g.Group(children),
	)
}