package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 9h7v9a20.41 20.41 0 0 1-6.5 1a9.67 9.67 0 0 1-4.64-1.03a6.8 6.8 0 0 1-2.89-3.04A11.21 11.21 0 0 1 2 10a9.9 9.9 0 0 1 1.1-4.78A7.7 7.7 0 0 1 6.35 2.1C7.77 1.37 9.95 1 12 1c.97 0 1.48.1 2.42.3c.93.19 1.81.37 2.58.7l-1 3c-.56-.28-1.5-.46-2.23-.64A9.24 9.24 0 0 0 11.53 4c-1.12 0-2.1.34-2.94.83A5.05 5.05 0 0 0 6.66 6.9A6.72 6.72 0 0 0 6 10a9 9 0 0 0 .48 3.09c.32.88.84 1.58 1.53 2.08c.7.5 1.62.83 2.75.83c.37 0 .8.02 1.08 0c.3-.03.66 0 1.16-.34V12h-3V9Z"/>`),
		g.Group(children),
	)
}