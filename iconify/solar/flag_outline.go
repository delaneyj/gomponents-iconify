package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 1.25a.75.75 0 0 1 .75.75v1.085l1.574-.315a9.427 9.427 0 0 1 5.35.492l.203.081a7.249 7.249 0 0 0 4.45.302a1.95 1.95 0 0 1 2.423 1.892v7.367c0 .988-.673 1.85-1.632 2.09l-.214.053a9.427 9.427 0 0 1-5.788-.393a7.927 7.927 0 0 0-4.498-.413l-1.868.374V22a.75.75 0 0 1-1.5 0V2A.75.75 0 0 1 5 1.25Zm.75 11.835l1.574-.315a9.428 9.428 0 0 1 5.35.492a7.927 7.927 0 0 0 4.866.33l.215-.054a.654.654 0 0 0 .495-.634V5.537a.45.45 0 0 0-.559-.437a8.749 8.749 0 0 1-5.371-.364l-.204-.082a7.927 7.927 0 0 0-4.498-.413l-1.868.374v8.47Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}