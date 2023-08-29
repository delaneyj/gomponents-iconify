package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttractionEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.5 1.5s-.5 0-.7.5l-.3.5H1c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h9c.6 0 1-.4 1-1v-5c0-.6-.4-1-1-1H7.5L7.2 2c-.2-.5-.7-.5-.7-.5h-2zm1 2C6.9 3.5 8 4.6 8 6S6.9 8.5 5.5 8.5S3 7.4 3 6s1.1-2.5 2.5-2.5zm0 1.5c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}