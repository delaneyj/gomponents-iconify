package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftTextTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 7.25A3.25 3.25 0 0 1 5.25 4h13.5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20H5.25A3.25 3.25 0 0 1 2 16.75v-9.5Zm8 11.25h8.75a1.75 1.75 0 0 0 1.75-1.75v-9.5a1.75 1.75 0 0 0-1.75-1.75H10v13Zm-5.25-10c0 .414.336.75.75.75h1.25a.75.75 0 0 0 0-1.5H5.5a.75.75 0 0 0-.75.75Zm0 3.5c0 .414.336.75.75.75h1.25a.75.75 0 0 0 0-1.5H5.5a.75.75 0 0 0-.75.75Zm.75 2.75a.75.75 0 0 0 0 1.5h1.25a.75.75 0 0 0 0-1.5H5.5Z"/>`),
		g.Group(children),
	)
}