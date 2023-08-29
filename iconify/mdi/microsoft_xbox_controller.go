package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftXboxController(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 15.75C6.75 15.75 6 18 4 19c-2 0-3.5-3 .5-11.5h.25l.44-.83S8 5 9.33 6.23h5.34c1.33-1.23 4.14.44 4.14.44l.44.83h.25C23.5 16 22 19 20 19c-2-1-2.75-3.25-4.75-3.25h-6.5M12 7a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}