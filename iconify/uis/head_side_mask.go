package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSideMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.2 2c-2-.1-4 .7-5.5 2.1c-1.4 1.4-2.2 3.3-2.2 5.4l-1.8 3c-.1.2-.2.3-.2.5c0 .1 0 .2.1.3L5 17.2v.1c.5 1 1.5 1.7 2.7 1.7h.8v2c0 .6.4 1 1 1s1-.4 1-1v-2h2c.1 0 .2 0 .3-.1l3.7-1.3v.1l1 3.5c.1.4.5.7 1 .7h.3c.5-.2.8-.7.7-1.2l-.9-3.2l1.9-7.3v-.4c0-4.1-3.2-7.6-7.3-7.8zM17 15.4l-3.5 1.2v-2.9l4.3-1.4l-.8 3.1zm1.4-5.4h-.3l-5.8 2h-6l1.1-1.7c.1-.2.2-.4.1-.6v-.2C7.5 6.5 10 4 13 4h.2c3 .2 5.4 2.7 5.3 5.8l-.1.2z"/>`),
		g.Group(children),
	)
}