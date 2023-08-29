package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1011.457 162l-416 807q-20 39-58.5 50.5t-72.5-10.5q-20-13-33-36q0-1-1-2t-1-2l-416-807q-20-39-9.5-82t44.5-65t72.5-10.5t58.5 49.5l333 646l333-646q20-38 58.5-49.5t72.5 10.5t44.5 65t-9.5 82zm-499 94q-53 0-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}