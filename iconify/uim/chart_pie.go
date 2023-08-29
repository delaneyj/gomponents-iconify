package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12V2c5.523 0 10 4.477 10 10H12z" opacity=".25"/><path fill="currentColor" d="m12 12l5 8.66A10.01 10.01 0 0 0 22 12H12z" opacity=".5"/><path fill="currentColor" d="M17 20.66L12 12V2c-5.523.002-9.999 4.48-9.997 10.003c.002 5.523 4.48 9.999 10.004 9.997A10 10 0 0 0 17 20.662l.003-.005l-.004.003z"/>`),
		g.Group(children),
	)
}