package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 99q-23 0-39 18l-68-68Q150 3 213 3q62 0 106 43.5T363 152q0 48-37 117l-77-78q18-16 18-39q0-22-16-37.5T213 99zm94 204l77 78l-27 27l-72-71q-16 23-34 46.5T223 418l-10 11q-6-6-16-18t-35.5-46.5t-45.5-67T80 224t-16-72q0-16 4-33L0 51l27-27l178 178l3 3z"/>`),
		g.Group(children),
	)
}