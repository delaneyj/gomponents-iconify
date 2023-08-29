package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bringtofront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 1024h-384q-27 0-45.5-18.5t-18.5-45.5v-64h64v64h384V576h-64v-64h64q26 0 45 19t19 45v384q0 27-19 45.5t-45 18.5zm-192-192h-512q-27 0-45.5-18.5t-18.5-45.5V256q0-26 18.5-45t45.5-19h512q27 0 45.5 19t18.5 45v512q0 27-19 45.5t-45 18.5zm-320-768h-384v384h64v64h-64q-27 0-45.5-18.5T.193 448V64q0-26 18.5-45t45.5-19h384q27 0 45.5 19t18.5 45v64h-64V64z"/>`),
		g.Group(children),
	)
}