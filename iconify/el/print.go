package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M367.456 73.572v181.2h-47.388L40.21 373.865h1119.58L879.932 254.772h-47.388v-181.2H367.456zM0 443.444v364.381h185.742l-65.259 318.603h959.033l-65.259-318.604H1200V443.443L0 443.444zm291.431 137.329h617.065l93.75 457.765H197.681l93.75-457.765z"/>`),
		g.Group(children),
	)
}