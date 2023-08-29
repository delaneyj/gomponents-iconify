package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UneditableTwoMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 2048l-633-158l-583-583l-723 722l-90-90L1939 19l90 90l-722 723l583 583l158 633zm-505-258l329 82l-82-329q-47 10-87 32t-73 55t-55 73t-32 87zm-327-867l-293 293l505 506q16-52 44-98t67-85t84-66t99-45l-506-505zM0 336q0-70 26-131T98 99t107-72T335 0q67 0 128 25t110 73l530 531l-90 90l-373-372l-293 293l372 373l-90 90L98 573q-48-48-73-109T0 336zm128 0q0 38 10 66t29 53t41 47t48 47l293-293q-25-25-47-48t-46-41t-54-28t-67-11q-43 0-80 16t-66 45t-44 66t-17 81z"/>`),
		g.Group(children),
	)
}