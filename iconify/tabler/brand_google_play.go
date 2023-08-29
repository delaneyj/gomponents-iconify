package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 3.71v16.58a.7.7 0 0 0 1.05.606l14.622-8.42a.55.55 0 0 0 0-.953L5.05 3.104A.7.7 0 0 0 4 3.711zM15 9L4.5 20.5m0-17L15 15"/>`),
		g.Group(children),
	)
}