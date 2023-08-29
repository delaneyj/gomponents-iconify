package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NodeJs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.245.203l-12.49 7.24a1.498 1.498 0 0 0-.75 1.313V23.24c0 .542.286 1.042.75 1.307l12.495 7.25a1.495 1.495 0 0 0 1.505 0l12.49-7.245c.464-.271.75-.771.75-1.307V8.755c0-.542-.286-1.042-.755-1.313L16.756.202a1.495 1.495 0 0 0-1.505 0z"/>`),
		g.Group(children),
	)
}