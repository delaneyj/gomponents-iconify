package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleIncludeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 10.75A4.75 4.75 0 0 1 10.75 6h12v16.75H6v-12Zm0 14.5v12A4.75 4.75 0 0 0 10.75 42h11.917A6.228 6.228 0 0 1 21 37.75v-10.5c0-.7.115-1.372.327-2H6Zm36-14.5v11.917A6.228 6.228 0 0 0 37.75 21h-10.5c-.7 0-1.372.115-2 .327V6h12A4.75 4.75 0 0 1 42 10.75Zm-19 16.5A4.25 4.25 0 0 1 27.25 23h10.5A4.25 4.25 0 0 1 42 27.25v10.5A4.25 4.25 0 0 1 37.75 42h-10.5A4.25 4.25 0 0 1 23 37.75v-10.5Z"/>`),
		g.Group(children),
	)
}