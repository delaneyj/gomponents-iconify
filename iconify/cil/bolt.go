package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m331.464 192l77-176H147.879l-81 288H187.9l-39.53 192h58.851l268.235-304ZM192.779 464h-5.149l39.529-192H109.121l63-224h187.415l-77 176h122.009Z"/>`),
		g.Group(children),
	)
}