package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emailalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.5 384L.5 125V64q0-27 19-45.5T64.5 0h896q26 0 45 18.5t19 45.5v61zm-32 64h64l480-249v505q0 26-19 45t-45 19h-896q-26 0-45-19t-19-45V199z"/>`),
		g.Group(children),
	)
}