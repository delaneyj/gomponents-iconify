package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2s1.343-2 3-2s3 .895 3 2Zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2s1.343-2 3-2s3 .895 3 2ZM9 10l12-3"/>`),
		g.Group(children),
	)
}