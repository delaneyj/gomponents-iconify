package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NuclearOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="256" r="192" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32"/><circle cx="256" cy="256" r="64" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32" d="M224 192L171 85m117 107l53-107m-13.45 170.81l119.41.13m-147.87 57.19l72.25 95.06M184.45 255.81l-119.41.13m147.87 57.19l-72.25 95.06"/>`),
		g.Group(children),
	)
}