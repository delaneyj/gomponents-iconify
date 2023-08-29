package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBellOffFill0"><g id="evaBellOffFill1"><path id="evaBellOffFill2" fill="currentColor" d="m15.88 18.71l-.59-.59L14 16.78l-.07-.07L6.58 9.4L5.31 8.14a5.68 5.68 0 0 0 0 .59v4.67l-1.8 1.81A1.64 1.64 0 0 0 4.64 18H8v.34A3.84 3.84 0 0 0 12 22a3.88 3.88 0 0 0 4-3.22ZM14 18.34A1.88 1.88 0 0 1 12 20a1.88 1.88 0 0 1-2-1.66V18h4ZM7.13 4.3l1.46 1.46l9.53 9.53l2 2l.31.3a1.58 1.58 0 0 0 .45-.6a1.62 1.62 0 0 0-.35-1.78l-1.8-1.81V8.94a6.86 6.86 0 0 0-5.83-6.88a6.71 6.71 0 0 0-5.32 1.61a6.88 6.88 0 0 0-.58.54Zm13.58 14.99L19.41 18l-2-2l-9.52-9.53L6.42 5L4.71 3.29a1 1 0 0 0-1.42 1.42L5.53 7l1.75 1.7l7.31 7.3l.07.07L16 17.41l.59.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}