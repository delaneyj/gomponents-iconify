package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 640V384h256v256H384zm512 0V384h256v256H896zm512-256h256v256h-256V384zM384 1152V896h256v256H384zm512 0V896h256v256H896zm512 0V896h256v256h-256zM384 1664v-256h256v256H384zm512 0v-256h256v256H896zm512 0v-256h256v256h-256z"/>`),
		g.Group(children),
	)
}