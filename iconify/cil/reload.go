package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M265.614 206.387H456V16h-32v133.887l-26.137-26.137c-79.539-79.539-208.96-79.54-288.5 0s-79.539 208.96 0 288.5a204.232 204.232 0 0 0 288.5 0l-22.627-22.627c-67.063 67.063-176.182 67.063-243.244 0s-67.063-176.183 0-243.246s176.182-67.063 243.245 0l28.01 28.01H265.614Z"/>`),
		g.Group(children),
	)
}