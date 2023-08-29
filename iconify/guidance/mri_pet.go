package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MriPet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.765 7A10.498 10.498 0 0 1 13 1.5c5.799 0 10.5 4.701 10.5 10.5S18.799 22.5 13 22.5A10.503 10.503 0 0 1 3.289 16M18 16H3.289M0 16h3.289M2 13h9.5v-1a3 3 0 0 0-3-3c-.54 0-4.692.437-8.4.84m13.4 1.56s1 1.6 2.25 1.6a1.75 1.75 0 0 0 1.75-1.75c0-.966-.784-1.746-1.75-1.746c-1.25 0-2.25 1.596-2.25 1.596v.3Z"/>`),
		g.Group(children),
	)
}