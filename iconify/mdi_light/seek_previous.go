package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeekPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.598 12.5L18 6v13l-1-.625L7.598 12.5Zm1.887 0L17 17.196V7.804L9.485 12.5ZM5 6h1v13H5V6Z"/>`),
		g.Group(children),
	)
}