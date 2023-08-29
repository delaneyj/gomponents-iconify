package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoisonCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 106c-33.81 0-61.887 22.69-71.25 53.438C174.532 154.22 163.258 151 151 151c-41.42 0-75 33.58-75 75c0 1.784.346 3.405.468 5.157C41.284 243.387 16 276.65 16 316c0 49.706 40.294 90 90 90h300c49.706 0 90-40.294 90-90c0-39.35-25.284-72.614-60.468-84.843c.123-1.752.468-3.374.468-5.157c0-41.42-33.58-75-75-75c-12.258 0-23.532 3.222-33.75 8.437C317.887 128.69 289.81 106 256 106zm-60 90l60 60l60-60l30 30l-60 60l60 60l-30 30l-60-60l-60 60l-30-30l60-60l-60-60z"/>`),
		g.Group(children),
	)
}