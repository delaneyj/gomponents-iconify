package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m512.426 449l-449-449h898zm-505 284q-7-14-7-29V64q0-15 7-29l350 349zm466-232q6 5 15.5 8t16.5 3h7q26 1 39-11l71-72l339 339h-898l339-339zm544-466q7 14 7 29v640q0 15-7 29l-350-349z"/>`),
		g.Group(children),
	)
}