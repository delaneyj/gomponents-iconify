package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryQuarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 8v16h24v-5h2v-6h-2V8zm2 2h20v12H5zm2 2v8h4v-8z"/>`),
		g.Group(children),
	)
}