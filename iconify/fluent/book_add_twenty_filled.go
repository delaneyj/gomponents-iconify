package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v5.207A5.5 5.5 0 0 0 9.207 16H5a1 1 0 0 0 1 1h3.6c.183.358.404.693.657 1H6a2 2 0 0 1-2-2V4Zm10.5 15a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-7a.5.5 0 0 1 .5.5V14h1.5a.5.5 0 0 1 0 1H15v1.5a.5.5 0 0 1-1 0V15h-1.5a.5.5 0 0 1 0-1H14v-1.5a.5.5 0 0 1 .5-.5ZM6.75 4a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75h6.5a.75.75 0 0 0 .75-.75v-.5a.75.75 0 0 0-.75-.75h-6.5Z"/>`),
		g.Group(children),
	)
}