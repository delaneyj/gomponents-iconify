package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowsGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 896q53 0 99-20t82-55t55-81t20-100q0-80 30-150t82-122t122-82t150-30v128q-53 0-99 20t-82 55t-55 81t-20 100q0 97-45 181T213 960q81 54 126 138t45 182q0 53 20 99t55 82t81 55t100 20v128q-80 0-150-30t-122-82t-82-122t-30-150q0-53-20-99t-55-82t-81-55t-100-20V896zm2048-640v640H768V256h1280zm-128 128H896v384h1024V384zM768 1024h1280v640H768v-640zm128 512h1024v-384H896v384z"/>`),
		g.Group(children),
	)
}