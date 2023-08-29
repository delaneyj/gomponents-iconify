package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KtorLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#F4F2ED" rx="60"/><path fill="#F26636" d="M95.406 34L31 96.46l64.406 62.428V96.46h64.374L95.406 34Z"/><path fill="#6373FF" d="m224.17 158.873l-64.406-62.429v62.429H95.39l64.374 62.459l64.406-62.459Z"/><path fill="#27282C" d="M159.78 96.444H95.39v62.444h64.39V96.444Z"/></g>`),
		g.Group(children),
	)
}