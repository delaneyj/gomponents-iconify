package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joystickatari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-19T0 960V832q0-26 19-45t45-19h256q0-26 18.5-45t45.5-19V192q0-20 20.5-42t43.5-22h192q27 0 45.5 19t18.5 45l-64 128v384q27 0 45.5 19t18.5 45h256q26 0 45 19t19 45v128q0 27-18.5 45.5T960 1024zM511.5 0Q538 0 557 19t19 45H448q0-26 18.5-45t45-19zM128 640h64q26 0 45 19t19 45H64q0-26 18.5-45t45.5-19z"/>`),
		g.Group(children),
	)
}