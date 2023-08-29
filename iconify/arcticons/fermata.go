package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fermata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 38.432V9.566m5.894-3.4l24.994 14.431m0 6.807l-24.994 14.43M7.45 9.566c0-3.4 2.952-5.098 5.894-3.4M7.45 38.432c0 3.398 2.95 5.101 5.894 3.402m24.994-14.43c2.95-1.703 2.95-5.103 0-6.806"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="20.966" cy="26.952" r="2.615"/><path d="M29.418 25.643c0-9.802-17-9.807-17 0c5.611-5.486 11.433-5.567 17 0Z"/></g>`),
		g.Group(children),
	)
}