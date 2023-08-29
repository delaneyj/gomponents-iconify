package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3H6zm2 2h7v1h2V5h7v22H8V5zm7 2v2h2V7h-2zm0 3v2h2v-2h-2zm0 3v2.188c-1.156.418-2 1.52-2 2.812c0 1.645 1.355 3 3 3s3-1.355 3-3c0-1.292-.844-2.394-2-2.813V13h-2zm1 4c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1z"/>`),
		g.Group(children),
	)
}