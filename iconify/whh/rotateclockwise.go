package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotateclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.488 1024h-576q-27 0-45.5-19t-18.5-45V832q0-27 18.5-45.5t45.5-18.5h576q27 0 45.5 18.5t18.5 45.5v128q0 26-18.5 45t-45.5 19zm-64-192h-448q-26 0-45 18.5t-19 45t19 45.5t45 19h448q27 0 45.5-19t18.5-45.5t-18.5-45t-45.5-18.5zm-38-138q-11 10-26 10t-26-10l-155-200q-11-10-11-23t11-23h109q-22-84-90.5-138t-157.5-54h-64q-26 0-45-19t-19-45.5t19-45t45-18.5h64q142 0 248.5 91.5t129.5 228.5h123q11 10 11 23t-11 23zm-666 10h-128q-27 0-45.5-19t-18.5-45V64q0-27 18.5-45.5T64.488 0h128q26 0 45 18.5t19 45.5v576q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}