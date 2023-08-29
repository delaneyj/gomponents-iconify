package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 576q0 221-147.5 384.5T640 1148v612q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-612q-217-24-364.5-187.5T0 576q0-117 45.5-223.5t123-184t184-123T576 0t223.5 45.5t184 123t123 184T1152 576zm-576 448q185 0 316.5-131.5T1024 576T892.5 259.5T576 128T259.5 259.5T128 576t131.5 316.5T576 1024z"/>`),
		g.Group(children),
	)
}