package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 4.6V17c0 1.9-.5 3-2.4 3H9.5c-.9 0-1.8-.4-2.4-1l-4.6-5l-.5-1c0-1 .5-1 .5-1c.3 0 .6 0 1 .2L7 14V3.3C7 2.6 7.3 2 8 2c.6 0 1 .7 1 1.4V9h1V1.2c0-.6.3-1.2 1-1.2s1 .6 1 1.3V9h1V2c0-.7.3-1.3 1-1.3s1 .6 1 1.3v7h1V4.6c0-.7.3-1.3 1-1.3s1 .6 1 1.3z"/>`),
		g.Group(children),
	)
}