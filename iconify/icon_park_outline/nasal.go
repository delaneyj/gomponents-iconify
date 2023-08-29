package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nasal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M16.988 4c.12 6.25-.203 10.586-.97 13.01c-1.152 3.633-9.936 13.753-9.936 19.092c0 5.34 6.304 8.287 9.709 6.277M32.002 4c-.225 6.25.047 10.586.814 13.01c1.15 3.633 10.143 12.44 10.143 18.282c0 5.843-6.512 9.097-9.917 7.087"/><path d="M13.006 34.834c1.71-.973 3.196-.973 4.46 0C19.36 36.292 19.956 40 24.008 40c4.053 0 6.04-4.157 7.992-5.166c1.302-.673 2.635-.673 4 0M20.285 22.145c-.794.855-1.634 1.808-2.247 3.469c-.465 1.259-.558 1.989-.622 2.81"/></g>`),
		g.Group(children),
	)
}