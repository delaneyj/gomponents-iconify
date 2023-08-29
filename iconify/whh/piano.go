package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.553 1024h-64q-27 0-45.5-19t-18.5-45V512q26 0 45-19t19-45V0h64q26 0 45 19t19 45v896q0 26-19 45t-45 19zm-192-512v448q0 26-19 45t-45 19h-64q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h64v448q0 27 18.5 45.5t45.5 18.5zm-320 512h-64q-27 0-45.5-19t-18.5-45V512q26 0 45-19t19-45V0h64q26 0 45 19t19 45v896q0 26-19 45t-45 19zm-256 0h-64q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h64v448q0 26 18.5 45t45.5 19v448q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}