package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UneditableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M633 1890L0 2048l158-633l584-583L19 109l90-90l1920 1920l-90 90l-723-722l-583 583zm-457-18l329-82q-10-46-32-87t-55-73t-73-54t-87-33l-82 329zm950-656L832 923l-506 505q52 17 98 45t85 66t66 84t45 99l506-506zm294-113l-91-90l373-373l-294-293l-372 372l-91-90l530-531h1q47-48 108-73t129-25q69 0 130 26t106 72t72 107t27 131q0 66-25 127t-73 110l-530 530zm439-621q29-29 45-67t16-79q0-43-16-81t-45-66t-66-44t-81-17q-38 0-66 10t-53 29t-47 41t-47 48l293 293l67-67h1h-1z"/>`),
		g.Group(children),
	)
}