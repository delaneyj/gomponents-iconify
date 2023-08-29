package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warmedalalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M766 384q-5 17-16 28L622 540l-45-12V0h127q27 0 45.5 19T768 64v320h-2zm-414-64l-96 192V0h257v512l-96-192h-65zM18 412Q7 401 2 384H1v-6q-2-10 0-21V64q0-26 18.5-45T65 0h127v528l-46 12zm270 187l96-214l96 214l225 41l-165 151l42 233l-198-114l-197 114l42-233L64 640z"/>`),
		g.Group(children),
	)
}