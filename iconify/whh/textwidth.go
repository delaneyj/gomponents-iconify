package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textwidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1016.553 950l-82 65q-8 9-19 9t-19-9v-54h-768v54q-8 9-19 9t-19-9l-82-65q-8-9-8-21.5t8-21.5l82-66q8-8 19-8t19 8v56h768v-56q8-8 19-8t19 8l82 66q8 9 8 21.5t-8 21.5zm-121-694q-26 0-45-18.5t-19-45.5t-49.5-45.5t-141.5-18.5q-27 0-45.5 19t-18.5 45v448q0 27 18.5 45.5t45 18.5t45.5 19t19 45.5t-19 45t-45 18.5h-256q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45-19t45.5-18.5t19-45.5V320v13v-141q0-26-19-45t-45-19q-92 0-142 18.5t-50 45.5t-19 45.5t-45.5 18.5t-45-18.5t-18.5-45.5V64q0-26 18.5-45t45.5-19h768q26 0 45 19t19 45v128q0 27-19 45.5t-45 18.5h-1z"/>`),
		g.Group(children),
	)
}