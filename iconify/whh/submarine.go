package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Submarine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 615q0 11-9 18t-23 7t-26.5-7t-18.5-18l-41-29q-43 16-124.5 75T684 730l-30 80q-4 9-16 15.5t-26 6.5H348q-14 0-26-6.5T306 810l-16-42h-34q-106 0-181-75T0 512q0-91 56.5-160.5T200 262l42-112q4-9 16-15.5t26-6.5h36V96q0-13-9.5-22.5T288 64h-32V0h64q26 0 45 19t19 45v64h36q14 0 26 6.5t16 15.5l40 106h74q36 0 80 22.5t80 51t82.5 60.5t81.5 45l47-90q6-11 18.5-18t26.5-7t23 7t9 18v270zM192.5 448q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19zm192 0q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19zm192 0q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19z"/>`),
		g.Group(children),
	)
}