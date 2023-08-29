package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Financial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 896h128v1152h-128V896zM512 1344l128-128v832H512v-704zm-256 256l128-128v576H256v-448zm512-512l128-128v1088H768v-960zm256-128l128 128v959h-128V960zm320 320l64-64v832h-128v-832l64 64zm192-192l128-128v1088h-128v-960zM0 1856l128-128v320H0v-192zM2048 256v512h-128V475l-576 575l-384-384L0 1627v-182l960-959l384 384l485-486h-293V256h512z"/>`),
		g.Group(children),
	)
}