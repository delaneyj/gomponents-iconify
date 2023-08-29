package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 12.2c-2.3 0-4.2-1.9-4.2-4.2s1.9-4.2 4.2-4.2s4.2 1.9 4.2 4.2c.1 2.3-1.9 4.2-4.2 4.2zm0-7C6 5.2 4.8 6.5 4.8 8s1.2 2.8 2.8 2.8s2.8-1.2 2.8-2.8S9 5.2 7.5 5.2z"/><path fill="currentColor" d="M8 16c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8c0 1.5-.4 3-1.2 4.2c-.3.5-1.1 1.2-2.3 1.2c-.8 0-1.3-.3-1.6-.6c-.7-.7-.6-1.8-.6-1.9V4h1.5v7c0 .2 0 .6.2.8c0 0 .2.2.5.2c.7 0 1.1-.5 1.1-.5c.6-1 1-2.2 1-3.4c0-3.6-2.9-6.5-6.5-6.5S1.5 4.4 1.5 8s2.9 6.5 6.5 6.5c.7 0 1.3-.1 1.9-.3l.4 1.4c-.7.3-1.5.4-2.3.4z"/>`),
		g.Group(children),
	)
}