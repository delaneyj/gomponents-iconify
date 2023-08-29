package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdInformation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M232 235h48v137h-48z" fill="currentColor"/><path d="M232 140h48v48h-48z" fill="currentColor"/>`),
		g.Group(children),
	)
}