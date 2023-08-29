package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M45 145v-38q0-8 7-15t15-7h170q5 0 9 2l34-29q-22-15-43-15H67q-28 0-46 18T3 107v170q0 26 18.5 45T67 341h170q27 0 45.5-19t18.5-45v-70l-42 38v32q0 8-7 15t-15 7H67q-8 0-15-7t-7-15V145zm314-43l-30-34l-28 28l-42 36l-81 71l-52-66l-34 25l77 105l90-79l42-37z"/>`),
		g.Group(children),
	)
}