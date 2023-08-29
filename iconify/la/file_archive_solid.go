package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileArchiveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h7v1h2V5h7v22H8zm7 2v2h2V7zm0 3v2h2v-2zm0 3v2.188c-1.156.417-2 1.519-2 2.812c0 1.645 1.355 3 3 3s3-1.355 3-3c0-1.293-.844-2.395-2-2.813V13zm1 4c.563 0 1 .438 1 1c0 .563-.438 1-1 1c-.563 0-1-.438-1-1c0-.563.438-1 1-1z"/>`),
		g.Group(children),
	)
}