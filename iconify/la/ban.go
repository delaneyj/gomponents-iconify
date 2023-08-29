package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.086 0 11 4.914 11 11c0 2.727-.988 5.207-2.625 7.125L9.031 7.469A10.95 10.95 0 0 1 16 5zM7.625 8.875l15.344 15.656A10.95 10.95 0 0 1 16 27C9.914 27 5 22.086 5 16c0-2.727.988-5.207 2.625-7.125z"/>`),
		g.Group(children),
	)
}