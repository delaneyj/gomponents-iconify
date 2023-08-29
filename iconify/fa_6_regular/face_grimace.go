package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrimace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48a208 208 0 1 0 0 416a208 208 0 1 0 0-416zm256 208a256 256 0 1 1-512 0a256 256 0 1 1 512 0zm-344 64c-13.3 0-24 10.7-24 24s10.7 24 24 24h8v-48h-8zm40 48h32v-48h-32v48zm96 0v-48h-32v48h32zm32 0h8c13.3 0 24-10.7 24-24s-10.7-24-24-24h-8v48zm-168-80h176c30.9 0 56 25.1 56 56s-25.1 56-56 56H168c-30.9 0-56-25.1-56-56s25.1-56 56-56zm-23.6-80a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}