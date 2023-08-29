package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlStackOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M17.24 19.399v-4.804h1.6V21H4.381v-6.405h1.598v4.804H17.24zM7.582 17.8h8.055v-1.604H7.582V17.8zm.195-3.64l7.859 1.641l.34-1.552l-7.861-1.642l-.338 1.553zm1.018-3.794l7.281 3.398l.678-1.463l-7.281-3.399l-.678 1.454v.01zm2.037-3.589l6.166 5.142l1.018-1.216l-6.162-5.14l-1.016 1.213l-.006.001zm3.982-3.778l-1.311.969l4.803 6.454l1.313-.971l-4.807-6.452h.002z" fill="currentColor"/>`),
		g.Group(children),
	)
}