package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path fill="currentColor" d="M10.373 16.589a.75.75 0 0 0 1.254.822l-1.254-.822Zm4.78-4.557a.75.75 0 0 0-1.255-.822l1.254.822ZM9.75 10A2.25 2.25 0 0 1 12 7.75v-1.5A3.75 3.75 0 0 0 8.25 10h1.5ZM12 7.75A2.25 2.25 0 0 1 14.25 10h1.5A3.75 3.75 0 0 0 12 6.25v1.5ZM14.25 10A2.25 2.25 0 0 1 12 12.25v1.5A3.75 3.75 0 0 0 15.75 10h-1.5ZM12 12.25A2.25 2.25 0 0 1 9.75 10h-1.5A3.75 3.75 0 0 0 12 13.75v-1.5Zm-.373 5.161l3.525-5.38l-1.254-.821l-3.525 5.379l1.254.822Z"/></g>`),
		g.Group(children),
	)
}