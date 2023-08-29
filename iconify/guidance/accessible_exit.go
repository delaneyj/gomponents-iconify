package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 23.5h5V.5h-15V8m8 15.5c-.727 0-1.273-1.071-1.273-1.071c-.545-1.072-.545-1.786-.545-2.858V18.5H8.374M7.5 12l-1.836 3.672M3 14l1.5-3s.8-.5 2.6-.5c2.4 0 3.9 2 3.9 2l-1.875 3.75m-3.461-.578a4 4 0 1 0-2.328 7.656a4 4 0 0 0 2.328-7.656Zm6.141-5.172s1.81-.557 2.135-1.776a1.768 1.768 0 0 0-1.242-2.163a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.962 2.61.962 2.61l.291.079Z"/>`),
		g.Group(children),
	)
}