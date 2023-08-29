package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v1152h-475l-256-256H0V640h1317l256-256h475zM128 1152h1152V768H128v384zm1499 256h37V512h-37l-219 219v458l219 219zm293-896h-128v896h128V512zm-832 384q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}