package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m15.925 14l-5.2 3.9c-.3.3-.8 0-.8-.5V2.6c0-.5.5-.8.8-.5l5.2 3.9m0 8h3c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1h-3m-9.5 8.5c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4c-1.6-1.6-1.6-4.1 0-5.7c.4-.4.3-1.1-.1-1.4c-.4-.3-.9-.3-1.3 0c-2.3 2.3-2.3 6.1 0 8.5c.1.2.4.3.7.3z"/>`),
		g.Group(children),
	)
}