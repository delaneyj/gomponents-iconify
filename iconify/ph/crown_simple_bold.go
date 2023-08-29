package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246.46 73.17a16 16 0 0 0-17.74-2.26l-46.9 23.38l-40-66.49a16.11 16.11 0 0 0-27.6 0l-40 66.49l-46.91-23.37A16.1 16.1 0 0 0 4.82 90.35l37 113.35a12 12 0 0 0 17.51 6.61C59.57 210.17 84.39 196 128 196s68.43 14.19 68.62 14.3a12 12 0 0 0 17.57-6.58l37-113.29a16 16 0 0 0-4.73-17.26Zm-50.93 110.35C182.18 178.4 159.2 172 128 172s-54.18 6.42-67.53 11.54l-27-82.71L70 119a16.19 16.19 0 0 0 21-6.11l37-61.49l37 61.5a16.18 16.18 0 0 0 21 6.1l36.52-18.2Z"/>`),
		g.Group(children),
	)
}