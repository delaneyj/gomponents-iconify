package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 5H8c-2.828 0-4.243 0-5.121.879C2 6.757 2 8.172 2 11v2c0 2.828 0 4.243.879 5.121C3.757 19 5.172 19 8 19h8c2.828 0 4.243 0 5.121-.879C22 17.243 22 15.828 22 13v-2c0-2.828 0-4.243-.879-5.121C20.243 5 18.828 5 16 5Z" opacity=".5"/><path d="M6 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-.25 3a.75.75 0 0 1-.75.75H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 .75.75Z"/></g>`),
		g.Group(children),
	)
}