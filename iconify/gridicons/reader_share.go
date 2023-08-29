package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<clipPath id="gridiconsReaderShare0"><path d="M0 0h20v20H0z"/></clipPath><g clip-path="url(#gridiconsReaderShare0)"><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.867 6.8V3l6.8 6.65l-6.8 6.65v-3.8s-10.2-.884-10.2 4.5c0-10.77 10.2-10.2 10.2-10.2z"/></g>`),
		g.Group(children),
	)
}