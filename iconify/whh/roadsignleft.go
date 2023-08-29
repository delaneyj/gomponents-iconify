package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roadsignleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 512H576v448q0 26-18.5 45t-45.5 19t-45.5-19t-18.5-45V512H256q-15 0-23-8q-9-3-12-7L15 355Q0 340 0 319.5T15 284l206-142q4-4 14-7q7-7 21-7h14q18-1 37 0h141V64q0-26 18.5-45T512 0t45.5 19T576 64v64h384q26 0 45 19t19 45v256q0 27-19 45.5T960 512z"/>`),
		g.Group(children),
	)
}