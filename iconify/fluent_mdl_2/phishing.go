package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phishing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v1280l-128-230V648l-896 447l-896-447v888h845l-64 128H0V384h2048zM1024 953l881-441H143l881 441zm384 455h128v320h-128v-320zm0 384h128v128h-128v-128zm64-896l576 1152H896l576-1152zm395 1040l-395-790l-395 790h790z"/>`),
		g.Group(children),
	)
}