package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M333 272L370 6h92l-70 266h-59zm20 39q-22 0-37 15t-15 36q0 22 15 37t37 15t37-15t15-37q0-21-15-36t-37-15zm-243-27h158l16-64H123l27-139h174l16-77H78L2 407h84z"/>`),
		g.Group(children),
	)
}