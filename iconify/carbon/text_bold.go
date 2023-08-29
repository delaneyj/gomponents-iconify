package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.25 25H9V7h8.5a5.25 5.25 0 0 1 4 8.65A5.25 5.25 0 0 1 18.25 25ZM12 22h6.23a2.25 2.25 0 1 0 0-4.5H12Zm0-7.5h5.5a2.25 2.25 0 1 0 0-4.5H12Z"/>`),
		g.Group(children),
	)
}