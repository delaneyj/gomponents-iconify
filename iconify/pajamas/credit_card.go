package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v1h13V4a.5.5 0 0 0-.5-.5ZM1.5 12V6.5h13V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5ZM2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H2Zm1.75 7.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}