package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openfolderalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M.232 576V64q0-27 18.5-45.5T64.232 0h384q27 0 45.5 18.5t18.5 45t18.5 45.5t45.5 19h256q27 0 45.5 18.5t18.5 45.5v64h-768zm1024-256l-192 512q0 26-18.5 45t-45.5 19h-704q-27 0-45.5-19t-18.5-45l192-512h832z"/>`),
		g.Group(children),
	)
}