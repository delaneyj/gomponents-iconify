package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrimace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm96-112h-8v-40h55.3c-3.8 22.7-23.6 40-47.3 40zm47.3-56H344v-40h8c23.8 0 43.5 17.3 47.3 40zm-71.3 0h-64v-40h64v40zm0 56h-64v-40h64v40zm-80-96v40h-64v-40h64zm0 56v40h-64v-40h64zm-80-16h-55.3c3.8-22.7 23.6-40 47.3-40h8v40zm0 56h-8c-23.8 0-43.5-17.3-47.3-40H168v40zm-23.6-192a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}