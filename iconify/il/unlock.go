package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M556 294v346q0 22-12 39t-32 26q-72 29-149 29H193q-77 0-149-29q-20-8-32-26T0 640V294h359V155q0-19-8-36t-21-27t-32-16t-37-1q-29 6-46 29t-18 53v32q0 12-12 12h-46q-5 0-8-4t-4-8v-34q0-34 14-63t37-50t54-31t66-6q28 4 53 17t41 35t27 47t10 56v134h127z"/>`),
		g.Group(children),
	)
}