package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15H3c-.6 0-1 .4-1 1s.4 1 1 1h4v4c0 .6.4 1 1 1s1-.4 1-1v-5c0-.6-.4-1-1-1zM8 2c-.6 0-1 .4-1 1v4H3c-.6 0-1 .4-1 1s.4 1 1 1h5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1zm8 7h5c.6 0 1-.4 1-1s-.4-1-1-1h-4V3c0-.6-.4-1-1-1s-1 .4-1 1v5c0 .6.4 1 1 1zm5 6h-5c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1s1-.4 1-1v-4h4c.6 0 1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}