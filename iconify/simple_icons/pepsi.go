package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pepsi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.43 3.277A10.839 10.839 0 0 0 2.718 17.594c7.455-2.033 13.503-7 15.712-14.317M12 22.84a10.839 10.839 0 0 0 9.21-16.574a7.607 7.607 0 0 1-2.873 8.195c-3.285 2.416-8.06 2.432-14.649 4.494A10.817 10.817 0 0 0 12 22.84M24 12A12 12 0 1 1 12 0a12 12 0 0 1 12 12"/>`),
		g.Group(children),
	)
}