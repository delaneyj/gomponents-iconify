package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnamusedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffdd67" d="M62 32c0 16.6-13.4 30-30 30S2 48.6 2 32S15.4 2 32 2s30 13.4 30 30z"/><path fill="#664e27" d="M53.9 29.2c-3.4-3.9-14-2.5-17.2 2.8c-.2.4.4 1 1.1 1.4c2.1-1.5 4.6-2.3 7.2-2.5c0 2.8 2.2 5.1 5 5.1c4 0 5.8-4.5 3.9-6.8m-27.4 0c-3.3-3.9-14-2.5-17.2 2.8c-.2.4.4 1 1.1 1.4c2.1-1.5 4.6-2.3 7.2-2.5c0 2.8 2.2 5.1 5 5.1c4 0 5.9-4.5 3.9-6.8m13.8 15.5c-5.8-1.5-12-.4-16.9 3c-1.2.9 1.1 4 2.3 3.2c3.2-2.3 8.4-3.8 13.7-2.4c1.3.3 2.4-3.4.9-3.8"/>`),
		g.Group(children),
	)
}