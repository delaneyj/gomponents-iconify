package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialYoutube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.8 8.6c-.2-1.5-.4-2.6-1-3C21.2 5.1 16 5 12 5s-9.2.1-9.8.6c-.6.4-.8 1.5-1 3S1 11 1 12s0 1.9.2 3.4s.4 2.6 1 3c.6.5 5.8.6 9.8.6c4 0 9.2-.1 9.8-.6c.6-.4.8-1.5 1-3S23 13 23 12s0-1.9-.2-3.4zm-12.8 7V8.4l6 3.6l-6 3.6z"/>`),
		g.Group(children),
	)
}