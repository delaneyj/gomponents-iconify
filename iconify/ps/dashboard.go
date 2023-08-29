package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m339 66l22-37Q305 8 256 8Q154 8 81 74T0 243q0 8 6 14.5t13 6.5h4q8 0 14-5.5t6-13.5q7-87 67.5-140.5T256 51q49 0 83 15zm98 83q27 45 30 96q0 7 5.5 13t13.5 6h5q8 0 14-7t5-14q-6-75-49-135zM448 8q-17-9-30 6l-21 35l-22 36l-96 162q-12-4-23-4q-35 0-60 25t-25 60t25 60t60 25t60-25t25-60q0-38-25-60l94-155l21-37l23-38q13-20-6-30zM256 371q-17 0-30-13t-13-30t13-30t30-13t30 13t13 30t-13 30t-30 13z"/>`),
		g.Group(children),
	)
}