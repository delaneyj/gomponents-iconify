package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M344 48h-24a16 16 0 0 0-16-16h-96a16 16 0 0 0-16 16h-24a56.16 56.16 0 0 0-56 56v247c0 35.3 144 65 144 65s144-29.7 144-65V104a56 56 0 0 0-56-56Zm-88 304a48 48 0 1 1 48-48a48 48 0 0 1-48 48Zm96-160a16 16 0 0 1-16 16H176a16 16 0 0 1-16-16v-64a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16ZM144 464h224m-32-32l48 48m-208-48l-48 48"/>`),
		g.Group(children),
	)
}