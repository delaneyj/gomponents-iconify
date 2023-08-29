package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 18.6L18.6 20l-4.5-4.5l-3.3 2.4c-.3.3-.8 0-.8-.5v-6L6.6 8c-.9 1.5-.7 3.3.5 4.7l.2.2c.4.4.4 1 0 1.4c-.2.2-.4.3-.7.3c-.3 0-.6-.1-.7-.3c-2-2.1-2.3-5.4-.7-7.7L3.7 5.1c-2.5 3.1-2.2 7.7.6 10.6c.2.2.3.4.3.7c0 .5-.4 1-1 1c-.2 0-.5-.1-.7-.3C-.7 13.4-1 7.6 2.3 3.7L0 1.4L1.4 0m9.4 2.1L16 6h3c.6 0 .9.3 1 .9V13c0 .6-.3.9-.9 1h-.8L10 5.8V2.7c0-.6.5-.9.8-.6z"/>`),
		g.Group(children),
	)
}