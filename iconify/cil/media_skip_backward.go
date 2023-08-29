package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaSkipBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M455.979 424.271A24.053 24.053 0 0 0 480 400.251V112.015a24 24 0 0 0-38.285-19.286L264 224.369V112.015a24 24 0 0 0-38.285-19.286L31.155 236.847a24 24 0 0 0 0 38.57l194.56 144.119A24 24 0 0 0 264 400.251V287.9l177.715 131.637a23.922 23.922 0 0 0 14.264 4.734ZM232 384.37L58.88 256.132L232 127.9ZM448 127.9v256.47L274.88 256.132Z"/>`),
		g.Group(children),
	)
}