package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Address(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.5 768h-896q-26 0-45-19t-19-45V64q0-27 19-45.5T64.5 0h896q26 0 45 18.5t19 45.5v640q0 26-19 45t-45 19zm-800-128h256q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-256q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5zm320-320h-320q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5zm0 128h-320q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5zm416-288q0-13-9.5-22.5t-22.5-9.5h-192q-13 0-22.5 9.5t-9.5 22.5v192q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5V160z"/>`),
		g.Group(children),
	)
}