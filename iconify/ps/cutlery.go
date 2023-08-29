package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M216 0h-21v459q0 23 15.5 38t37.5 15t37.5-15t15.5-38V282q22-19 22-47V107q0-45-31-76T216 0zm43 459q0 10-11 10t-11-10V299h22v160zm21-224q0 9-6 15t-15 6h-22V47q19 6 31 22.5t12 37.5v128zM3 288v171q0 23 15.5 38T56 512t37.5-15t15.5-38V288q0-22-15.5-37.5T56 235t-37.5 15.5T3 288zm42 0q0-11 11-11t11 11v171q0 10-11 10t-11-10V288zm43-139H67V0H45v149H24V0H3v171h42v64h22v-64h42V0H88v149z"/>`),
		g.Group(children),
	)
}