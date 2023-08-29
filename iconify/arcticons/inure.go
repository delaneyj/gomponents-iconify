package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.462c-6.73 0-17.448-12.79-17.448-25.91c0-9 7.523-11.062 17.448-11.062s17.448 2.062 17.448 11.062c0 13.121-10.718 25.91-17.448 25.91Zm-5.266-20.309L28.328 5.65M16.824 38.965l15.16-26.077"/>`),
		g.Group(children),
	)
}