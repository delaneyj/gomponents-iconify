package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18a3.996 3.996 0 0 0-3.858 3H17V11h5.142a4 4 0 1 0 0-2H17a2.002 2.002 0 0 0-2 2v4H9.858a4 4 0 1 0 0 2H15v4a2.002 2.002 0 0 0 2 2h5.142A3.993 3.993 0 1 0 26 18Zm0-10a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2ZM6 18a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm20 6a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}