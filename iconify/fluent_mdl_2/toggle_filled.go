package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1582 1490H466q-96 0-181-37t-148-100t-100-148t-37-181q0-96 37-181t100-148t148-100t181-37h1116q96 0 181 37t148 100t100 148t37 181q0 96-37 181t-100 148t-148 100t-181 37z"/>`),
		g.Group(children),
	)
}