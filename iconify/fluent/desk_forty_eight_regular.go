package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.25 8A4.25 4.25 0 0 0 4 12.25v21.5A6.25 6.25 0 0 0 10.25 40h7.5A6.25 6.25 0 0 0 24 33.75V18.5h17.5v20.25a1.25 1.25 0 1 0 2.5 0v-26.5A4.25 4.25 0 0 0 39.75 8H8.25ZM6.5 18.5h15v15.25a3.75 3.75 0 0 1-3.75 3.75h-7.5a3.75 3.75 0 0 1-3.75-3.75V18.5Zm0-2.5v-3.75c0-.966.784-1.75 1.75-1.75h31.5c.967 0 1.75.784 1.75 1.75V16h-35Zm4.75 7a1.25 1.25 0 1 0 0 2.5h5.5a1.25 1.25 0 1 0 0-2.5h-5.5Z"/>`),
		g.Group(children),
	)
}