package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 14.5V6H9.75A1.75 1.75 0 0 1 8 4.25V1.5H3.5v13h9Zm-.121-10L9.5 1.621V4.25c0 .138.112.25.25.25h2.629ZM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V1Zm5.75 10.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5ZM6.28 6.22a.75.75 0 0 0-1.06 1.06L6.44 8.5L5.22 9.72a.75.75 0 1 0 1.06 1.06l1.75-1.75l.53-.53l-.53-.53l-1.75-1.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}