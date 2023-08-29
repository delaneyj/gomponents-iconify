package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyFlowerPlantTreeFlowerPetalsBloom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="7" cy="7.26" rx="1.5" ry="1.49"/><path d="M13.37 5.34a2.57 2.57 0 0 0-3.24-1.64a2.3 2.3 0 0 0-.65.3a2.71 2.71 0 0 0 .1-.71a2.58 2.58 0 0 0-5.16 0a2.71 2.71 0 0 0 .1.71a2.3 2.3 0 0 0-.65-.31a2.56 2.56 0 1 0-1.59 4.87a3 3 0 0 0 .72.12a3 3 0 0 0-.5.51a2.54 2.54 0 0 0 .57 3.57a2.59 2.59 0 0 0 3.6-.56a2.47 2.47 0 0 0 .33-.64a2.47 2.47 0 0 0 .34.64a2.59 2.59 0 0 0 3.6.56a2.54 2.54 0 0 0 .57-3.57a3 3 0 0 0-.5-.51a3 3 0 0 0 .71-.12a2.55 2.55 0 0 0 1.65-3.22Z"/></g>`),
		g.Group(children),
	)
}