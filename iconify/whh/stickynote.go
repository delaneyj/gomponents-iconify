package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stickynote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1023 576q0 48-34 82L658 989q-34 34-82 34v1H0V0h1024v576h-1zm-447 0v378q26-9 44-27l307-307q18-18 27-44H576z"/>`),
		g.Group(children),
	)
}