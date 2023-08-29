package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sendbackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 1024h-512q-26 0-45-18.5t-19-45.5V640h-320q-27 0-45.5-18.5T.193 576V64q0-26 18.5-45t45.5-19h512q26 0 45 19t19 45v320h320q26 0 45 19t19 45v512q0 27-19 45.5t-45 18.5zm0-576h-512v512h512V448z"/>`),
		g.Group(children),
	)
}