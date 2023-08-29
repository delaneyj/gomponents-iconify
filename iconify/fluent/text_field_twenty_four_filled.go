package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFieldTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A3.75 3.75 0 0 1 5.75 3h12.5A3.75 3.75 0 0 1 22 6.75v10.5A3.75 3.75 0 0 1 18.25 21H5.75A3.75 3.75 0 0 1 2 17.25V6.75Zm10.75.75h2.75v.75a.75.75 0 0 0 1.5 0v-1.5a.75.75 0 0 0-.75-.75h-8.5a.75.75 0 0 0-.75.75v1.5a.75.75 0 1 0 1.5 0V7.5h2.75v9h-.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-.5v-9Z"/>`),
		g.Group(children),
	)
}