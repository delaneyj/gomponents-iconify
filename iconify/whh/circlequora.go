package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlequora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm256-544q0-119-75-203.5T512 192t-181 84.5T256 480t75 203.5T512 768q63 0 118-33q21 14 52 23.5t54 9.5q13 0 22.5-9.5T768 736t-9.5-22.5T736 704q-25 0-44-19q76-85 76-205zm-224 96q-13 0-22.5 9.5T512 608t9.5 22.5T544 640q24 0 44 19q-34 45-76 45q-53 0-90.5-65.5T384 480t37.5-158.5T512 256t90.5 65.5T640 480q0 60-18 114q-42-18-78-18z"/>`),
		g.Group(children),
	)
}