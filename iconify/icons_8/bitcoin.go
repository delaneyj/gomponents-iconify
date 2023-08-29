package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 3v3H8v20h4v3h2v-3h2v3h2v-3h1.5c3.026 0 5.5-2.474 5.5-5.5a5.52 5.52 0 0 0-2.875-4.844A5.517 5.517 0 0 0 24 11.5C24 8.474 21.526 6 18.5 6H18V3h-2v3h-2V3h-2zm-2 5h8.5c1.944 0 3.5 1.556 3.5 3.5S20.444 15 18.5 15H10V8zm0 9h9.5c1.944 0 3.5 1.556 3.5 3.5S21.444 24 19.5 24H10v-7z"/>`),
		g.Group(children),
	)
}