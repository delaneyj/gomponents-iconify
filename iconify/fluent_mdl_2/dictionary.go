package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 0v2048H256q-53 0-99-20t-82-55t-55-81t-20-100V256q0-49 21-95t57-82t82-57t96-22h1408zm-128 1664H256q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10h1280v-256zm-1408-94q60-34 128-34h1280V128H256q-23 0-46 11t-41 30t-29 41t-12 46v1314zM384 384h896v384H384V384zm128 256h640V512H512v128z"/>`),
		g.Group(children),
	)
}