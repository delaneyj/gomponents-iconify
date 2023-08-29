package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.47 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8s8 3.58 8 8s-3.58 8-8 8m-4-6L7 8l3 2l2-3l2 3l3-2l-1 6H8m.56 2c-.34 0-.56-.22-.56-.56V15h8v.44c0 .34-.22.56-.56.56H8.56Z"/>`),
		g.Group(children),
	)
}