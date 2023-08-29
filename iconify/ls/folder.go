package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M380 146h305c17 0 32 15 32 32v423c0 17-15 31-32 31H32c-17 0-32-14-32-31V178c0-17 15-32 32-32h34l16-53c0-17 15-32 32-32h218c17 0 32 15 32 32z"/>`),
		g.Group(children),
	)
}