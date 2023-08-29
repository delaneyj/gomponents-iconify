package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialInstagramCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9zm0 7c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2zm2.8-2c0-.7.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2s-1.2-.5-1.2-1.2zM12 19c-3.9 0-7-3.1-7-7h3c0 2.2 1.8 4 4 4s4-1.8 4-4h3c0 3.9-3.1 7-7 7z"/>`),
		g.Group(children),
	)
}