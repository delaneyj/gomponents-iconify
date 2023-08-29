package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M440 356q-10 0-16-7t-7-16V11q66 8 124 38t100 75t69 105t30 127H440zm-93 45q0 10 7 17t17 7h367q-9 72-44 133t-88 106t-121 67t-145 17q-68-6-128-35t-105-76t-74-106T1 403q-5-78 20-148t73-123t113-86T347 9v392z"/>`),
		g.Group(children),
	)
}