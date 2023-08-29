package fa_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 0H48C21.49 0 0 21.49 0 48v464l192-112l192 112V48c0-26.51-21.49-48-48-48zm0 428.43l-144-84l-144 84V54a6 6 0 0 1 6-6h276c3.314 0 6 2.683 6 5.996V428.43z"/>`),
		g.Group(children),
	)
}