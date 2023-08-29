package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1504 128q113 0 212 43t173 116t116 173t43 212q0 109-41 209t-118 176l-865 864l-865-864Q83 981 42 881T0 672q0-112 42-211t117-173t173-117t212-43q83 0 148 19t120 52t106 81t106 103q55-56 105-103t106-80t121-53t148-19zm294 838q59-59 90-135t31-159q0-87-32-162t-88-131t-132-87t-163-32q-84 0-149 26t-120 70t-105 97t-106 111q-54-54-105-109t-106-99t-121-72t-148-28q-86 0-162 32t-132 89t-89 133t-33 162q0 83 31 159t91 135l774 774l774-774z"/>`),
		g.Group(children),
	)
}