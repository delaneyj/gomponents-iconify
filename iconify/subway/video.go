package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M302.5 162.9h-256C20.9 162.9 0 183.8 0 209.5v256C0 491.1 20.9 512 46.5 512h256c25.7 0 46.5-20.9 46.5-46.5v-256c.1-25.7-20.8-46.6-46.5-46.6zm69.9 116.4v116.4L512 465.5v-256l-139.6 69.8zM69.8 139.6c38.6 0 69.8-31.3 69.8-69.8S108.4 0 69.8 0C31.3 0 0 31.3 0 69.8s31.3 69.8 69.8 69.8zM279.3 0c-38.6 0-69.8 31.3-69.8 69.8s31.3 69.8 69.8 69.8c38.6 0 69.8-31.3 69.8-69.8S317.8 0 279.3 0z"/>`),
		g.Group(children),
	)
}