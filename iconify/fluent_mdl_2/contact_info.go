package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M740 1077q65 33 117 81t90 108t57 128t20 142H896q0-79-30-149t-83-122t-122-82t-149-31q-79 0-149 30t-122 83t-82 122t-31 149H0q0-73 20-141t57-128t89-108t118-82q-73-54-114-136t-42-173q0-79 30-149t83-122t122-82t149-31q79 0 149 30t122 83t82 122t31 149q0 91-41 173t-115 136zM256 768q0 53 20 99t55 81t81 55t100 21q52 0 99-20t81-55t55-81t21-100q0-52-20-99t-55-81t-82-55t-99-21q-53 0-99 20t-81 55t-55 82t-21 99zm1792-384v128h-896V384h896zm-896 512h896v128h-896V896zm0 512h896v128h-896v-128z"/>`),
		g.Group(children),
	)
}