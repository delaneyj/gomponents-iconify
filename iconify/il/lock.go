package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M556 294v346q0 22-12 39t-32 26q-72 29-149 29H193q-77 0-149-29q-20-8-32-26T0 640V294h127V155q0-34 14-63t37-50t54-31t66-6q28 4 53 17t41 35t27 47t10 56v134h127zm-359 0h162V155q0-34-23-58t-58-24t-57 24t-24 58v139z"/>`),
		g.Group(children),
	)
}