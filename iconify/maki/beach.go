package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m5.36 1.67l-.01 4.02a4.452 4.452 0 0 0-1.1-.11c-.37.1-.74.63-1.1.76a4.202 4.202 0 0 1 2.21-4.67Zm2.41-.64L9.8 4.48a3.183 3.183 0 0 1 .84-.61c.36-.1.94.17 1.34.11a4.202 4.202 0 0 0-4.21-2.95ZM1 13h13c-.66-.66-2.64-1.11-4.34-1.33l-1.87-7c.52-.05 1.15.03 1.53 0l-2.11-3.6H7.2a6.174 6.174 0 0 0-.7.14a4.38 4.38 0 0 0-.64.22l-.01 4.15c.35-.17.84-.54 1.3-.74l1.8 6.74c-.58-.05-1.09-.08-1.45-.08C6.03 11.5 2 12 1 13Z"/>`),
		g.Group(children),
	)
}