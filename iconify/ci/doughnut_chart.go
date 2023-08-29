package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoughnutChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.051 21.949a10 10 0 0 1-1-19.949v4.04a5.994 5.994 0 1 0 6.91 6.909H22a10.015 10.015 0 0 1-9.95 9Zm9.95-11h-4.04a5.993 5.993 0 0 0-4.91-4.909V2a10.022 10.022 0 0 1 8.95 8.948v.001Z"/>`),
		g.Group(children),
	)
}