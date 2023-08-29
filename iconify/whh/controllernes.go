package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Controllernes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 568H64q-26 0-45-19T0 504V120q0-26 19-45t45-19h896q26 0 45 19t19 45v384q0 27-18.5 45.5T960 568zM384 248h-64v-64q0-26-18.5-45t-45-19t-45.5 19t-19 45v64h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64v64q0 27 19 45.5t45.5 18.5t45-18.5T320 440v-64h64q27 0 45.5-18.5t18.5-45t-18.5-45.5t-45.5-19zm320.5 64q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19zm192 0q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19z"/>`),
		g.Group(children),
	)
}