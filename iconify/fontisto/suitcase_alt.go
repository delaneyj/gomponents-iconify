package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.109 9.267h-2.281v-.689h-.998V1.164h1.259V0H4.4v1.164h1.255v7.414h-.998v.689H2.375a2.383 2.383 0 0 0-2.376 2.376v8.72a2.383 2.383 0 0 0 2.376 2.376h.689a1.26 1.26 0 1 0 2.52 0h8.294a1.26 1.26 0 1 0 2.52 0h.713a2.383 2.383 0 0 0 2.376-2.376v-8.72a2.383 2.383 0 0 0-2.376-2.376zM6.059 1.164h7.366v7.414h-.998v.689h-5.37v-.689h-.998zm-2.281 18.06L2 17.251l5.251-4.729l1.782 1.972zm7.272-8.126h2.851v2.662H11.05zm3.612 9.338a2.662 2.662 0 1 1 2.662-2.662v.005a2.658 2.658 0 0 1-2.657 2.657h-.005z"/>`),
		g.Group(children),
	)
}