package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditPhoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 576q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45zM128 1664h732l-32 128H0V256h2048v580q-28-22-60-37t-68-23V384H128v677l448-447l384 384l256-256l261 260l-91 91l-170-171l-166 166l171 170l-91 91l-554-555l-448 449v421zm1792-178l128-128v434h-434l128-128h178v-178zm128-392q0 39-15 76t-43 65l-717 717l-377 94l94-376l717-717q28-28 65-41t76-13q42 0 78 14t64 41t42 61t16 79zm-149 51q21-21 21-51q0-32-20-50t-52-19q-14 0-27 4t-23 14l-692 692l-34 135l135-34l692-691z"/>`),
		g.Group(children),
	)
}