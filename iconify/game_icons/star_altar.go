package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarAltar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 25.61l-97.8 36.67L256 86.72l97.8-24.44L256 25.61zM136.1 75.3L105.9 407h300.2L375.9 75.3l-119.9 30l-119.9-30zM256 128l30 72.7l78.3 6l-59.8 51L323 334l-67-41.2l-66.9 41.3l18.4-76.5l-59.8-50.9l78.4-6L256 128zM70.09 425l-24.8 62H237.3l-49.6-62H70.09zm140.61 0l45.3 56.6l45.3-56.6h-90.6zm113.6 0l-49.6 62h192l-24.8-62H324.3z"/>`),
		g.Group(children),
	)
}