package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Numbersign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M555 232v60H435l-24 161h114v60H402l-36 246h-60l35-246H180l-37 246H82l37-246H0v-60h128l24-161H31v-60h130L195 0h60l-33 232h161L418 0h61l-35 232h111zM189 453h162l23-161H212z"/>`),
		g.Group(children),
	)
}