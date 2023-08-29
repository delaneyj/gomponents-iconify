package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewOriginal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 896H256V256h1024v640zm-128-512H384v384h768V384zm896 256v1024H896v-128h1024V768h-384v384H768v256l-128-128v-128H0V0h1536v640h512zm-640-512H128v896h1280V128zM464 1414l315 314l-319 318l-74-74l179-180H384q-80 0-149-30t-122-82t-83-122t-30-150h128q0 53 20 99t55 82t81 55t100 20h181l-175-176l74-74z"/>`),
		g.Group(children),
	)
}