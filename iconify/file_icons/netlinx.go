package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netlinx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M453.676 283.287L374.094 78.443H237.68L0 433.389c95.384-50.707 239.414-106.51 453.676-150.102zm-164.623-84.586c4.08 11.557 27.014 76.667 30.955 87.82H233.08c1.373-2.17 49.216-77.232 55.973-87.82zm83.052 234.856s-24.225-68.468-30.592-86.476c39.547-8.518 81.996-16.669 127.513-24.299L512 433.556l-139.895.001z"/>`),
		g.Group(children),
	)
}