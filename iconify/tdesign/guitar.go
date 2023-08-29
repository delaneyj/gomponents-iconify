package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guitar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 1.586L22.414 5L21 6.414l-1-1l-.586.586l.5.5L18.5 7.914l-.5-.5l-.828.829a5 5 0 0 1-2.183 7.346a7 7 0 1 1-6.577-6.577a5 5 0 0 1 7.345-2.183L16.586 6l-.5-.5L17.5 4.086l.5.5l.586-.586l-1-1L19 1.586ZM13 8a3.001 3.001 0 0 0-2.924 2.325l-.215.934l-.942-.175a5 5 0 1 0 3.997 3.997l-.175-.942l.934-.215A3.002 3.002 0 0 0 13 8Zm-5 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}