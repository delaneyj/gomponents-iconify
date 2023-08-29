package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisBallAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.66 7c2.762 4.783 1.123 10.899-3.66 13.66C12.217 23.422 6.101 21.783 3.34 17C.578 12.217 2.217 6.1 7 3.34C11.783.578 17.899 2.217 20.66 7Z"/><path d="M21.46 15.242c-4.986-3.303-7.582-7.8-7.538-13.056m-3.844 19.628C9.71 15.844 7.114 11.347 2.54 8.758"/></g>`),
		g.Group(children),
	)
}