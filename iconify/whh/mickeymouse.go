package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mickeymouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M869 436q27 68 27 140q0 104-51.5 192.5t-140 140T512 960t-192.5-51.5t-140-140T128 576q0-72 27-140q-69-22-112-80.5T0 224q0-93 65.5-158.5T224 0q86 0 149 57t72 141q36-6 67-6t67 6q9-84 72-141T800 0q93 0 158.5 65.5T1024 224q0 73-43 131.5T869 436z"/>`),
		g.Group(children),
	)
}