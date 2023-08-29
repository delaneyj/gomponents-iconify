package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tnt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#ff4081"/><path fill="#fff" d="M13.89 21.503L14.048 25s1.875-.318 3.828 0l.156-3.497zM11.626 18.8s5.469-.477 8.672 0l.312-3.815a59.109 59.109 0 0 0-9.14 0zM26 8.944S15.531 6.718 6 9.103l.703 4.371l.86-1.669s8.984-1.351 17.03 0l.782 1.67z"/></g>`),
		g.Group(children),
	)
}