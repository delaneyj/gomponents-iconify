package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1.19 0C1.08 0 1 .08 1 .19v7.63c0 .11.08.19.19.19h4.63c.11 0 .19-.08.19-.19V.19c0-.11-.08-.19-.19-.19H1.19zM2 1h3v5H2V1zm1.5 5.5c.28 0 .5.22.5.5s-.22.5-.5.5S3 7.28 3 7s.22-.5.5-.5z"/>`),
		g.Group(children),
	)
}