package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><circle cx="22.3" cy="31.6" r="5" fill="#664e27"/><path fill="#917524" d="M51.2 27.5c-3.2-2.7-7.5-3.9-11.7-3.1c-.6.1-1.1-2-.4-2.2c4.8-.9 9.8.5 13.5 3.6c.6.5-1 2.1-1.4 1.7m-26.7-8.7c-4.2-.7-8.5.4-11.7 3.1c-.4.4-2-1.2-1.4-1.7c3.7-3.2 8.7-4.5 13.5-3.6c.7.2.2 2.3-.4 2.2"/><path fill="#664e27" d="M50.2 34.3c-1.7-3.5-4.4-5.3-7-5.3s-5.2 1.8-7 5.3c-.2.4.7 1 1.2.6c1.7-1.3 3.7-1.8 5.8-1.8s4.1.5 5.8 1.8c.4.3 1.3-.3 1.2-.6m-6.1 7.9c-6.9 3.6-16.4 2.9-19.1 2.9c-.7 0-1.2.3-1 .9c2 7 17 7 21.1-2.7c.5-1.1-.3-1.4-1-1.1"/>`),
		g.Group(children),
	)
}