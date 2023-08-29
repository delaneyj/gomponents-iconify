package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.158 768h-384q-26 0-45-18.5t-19-45.5V64q0-26 19-45t45-19h384q26 0 45 19t19 45v640q0 27-18.5 45.5t-45.5 18.5zm-64-640h-256v512h256V128zm-512 640h-320q-26 0-45-18.5t-19-45.5t19-45.5t45-18.5h64V128h-64q-26 0-45-18.5t-19-45t19-45.5t45-19h128q27 0 45.5 19t18.5 45v576h64v-64q0-27 19-45.5t45.5-18.5t45 19t18.5 45v128q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}