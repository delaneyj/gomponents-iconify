package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 5h8c2.828 0 4.243 0 5.121.879C22 6.757 22 8.172 22 11v2c0 2.828 0 4.243-.879 5.121C20.243 19 18.828 19 16 19H8c-2.828 0-4.243 0-5.121-.879C2 17.243 2 15.828 2 13v-2c0-2.828 0-4.243.879-5.121C3.757 5 5.172 5 8 5Zm-2 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-.25 3a.75.75 0 0 1-.75.75H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 .75.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}