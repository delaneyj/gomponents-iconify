package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M512 1088q66 0 128-15v655q0 26-19 45t-45 19H448q-26 0-45-19t-19-45v-655q62 15 128 15zM512 0q212 0 362 150t150 362t-150 362t-362 150t-362-150T0 512t150-362T512 0zm0 224q14 0 23-9t9-23t-9-23t-23-9q-146 0-249 103T160 512q0 14 9 23t23 9t23-9t9-23q0-119 84.5-203.5T512 224z"/>`),
		g.Group(children),
	)
}