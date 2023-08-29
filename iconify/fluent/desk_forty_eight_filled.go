package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.25 8A4.25 4.25 0 0 0 4 12.25V16h37.5v22.75a1.25 1.25 0 1 0 2.5 0v-26.5A4.25 4.25 0 0 0 39.75 8H8.25ZM24 18.5H4v15.25A6.25 6.25 0 0 0 10.25 40h7.5A6.25 6.25 0 0 0 24 33.75V18.5Zm-14 6.75c0-.69.56-1.25 1.25-1.25h5.5a1.25 1.25 0 1 1 0 2.5h-5.5c-.69 0-1.25-.56-1.25-1.25Z"/>`),
		g.Group(children),
	)
}