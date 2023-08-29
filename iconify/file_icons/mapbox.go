package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mapbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M30.118 0v512h75.294V90.354l150.588 256l150.588-256V512h75.294V0h-60.235c-30.118-.096-53.772 7.53-67.059 30.118L256 197.648l-98.588-167.53C144.125 7.53 120.47-.096 90.353.001H30.118z"/>`),
		g.Group(children),
	)
}