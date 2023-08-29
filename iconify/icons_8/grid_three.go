package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6v20h24V6H4zm2 2h5v4H6V8zm7 0h13v4H13V8zm-7 6h5v4H6v-4zm7 0h13v4H13v-4zm-7 6h5v4H6v-4zm7 0h13v4H13v-4z"/>`),
		g.Group(children),
	)
}