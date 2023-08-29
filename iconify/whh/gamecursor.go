package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamecursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H704v256q0 26-19 45t-45 19H384q-27 0-45.5-19T320 960V704H64q-27 0-45.5-19T0 640V384q0-27 19-45.5T64 320h256V64q0-27 19-45.5T384 0h256q27 0 45.5 18.5T704 64v256h256q27 0 45.5 18.5T1024 384v256q0 26-19 45t-45 19zM192 384L64 512l128 128V384zM512 64L384 192h256zm-64 447.5q0 26.5 19 45.5t45.5 19t45-19t18.5-45.5t-19-45t-45-18.5t-45 18.5t-19 45zM384 832l128 128l128-128H384zm448-448v256l128-128z"/>`),
		g.Group(children),
	)
}