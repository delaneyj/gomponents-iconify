package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 10c-5.738 0-10.45 4.393-10.955 10h-.295a7.75 7.75 0 0 0 0 15.5h22.5a7.75 7.75 0 0 0 0-15.5h-.295C34.45 14.393 29.738 10 24 10Z"/>`),
		g.Group(children),
	)
}