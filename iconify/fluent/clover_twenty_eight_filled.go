package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloverTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.75 2a5.75 5.75 0 1 0 0 11.5h5a.75.75 0 0 0 .75-.75v-5A5.75 5.75 0 0 0 7.75 2Zm0 24a5.75 5.75 0 0 1 0-11.5h5a.75.75 0 0 1 .75.75v5A5.75 5.75 0 0 1 7.75 26Zm12.5-24a5.75 5.75 0 0 1 0 11.5h-5a.75.75 0 0 1-.75-.75v-5A5.75 5.75 0 0 1 20.25 2Zm0 24a5.75 5.75 0 0 0 0-11.5h-5a.75.75 0 0 0-.75.75v5A5.75 5.75 0 0 0 20.25 26Z"/>`),
		g.Group(children),
	)
}