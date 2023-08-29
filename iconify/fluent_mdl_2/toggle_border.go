package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleBorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1582 558q96 0 181 37t148 100t100 148t37 181q0 96-37 181t-100 148t-148 100t-181 37H466q-96 0-181-37t-148-100t-100-148t-37-181q0-96 37-181t100-148t148-100t181-37h1116zm0 838q77 0 145-29t118-80t80-118t30-145q0-77-29-144t-80-118t-119-80t-145-30H466q-77 0-145 29t-118 80t-80 118t-30 145q0 77 29 144t80 118t119 80t145 30h1116z"/>`),
		g.Group(children),
	)
}