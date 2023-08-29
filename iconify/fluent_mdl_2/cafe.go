package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1663 512q79 0 149 30t122 82t83 123t30 149q0 80-30 149t-82 122t-123 83t-149 30h-103q-44 77-105 142t-135 114h279q26 0 45 19t19 45q0 26-19 45t-45 19H191q-26 0-45-19t-19-45q0-26 19-45t45-19h279q-81-54-144-124t-108-152t-67-175t-24-189V512h1536zM895 1536q88 0 170-23t153-64t129-100t100-130t65-153t23-170V640H255v256q0 88 23 170t64 153t100 129t130 100t153 65t170 23zm768-384q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q0 65 1 130t-1 129t-12 127t-32 126h44z"/>`),
		g.Group(children),
	)
}