package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m14.57-1.398a2 2 0 0 0 2.83 2.827m-9.287-9.296a3 3 0 1 0 3.765 3.715M12 7v1m-5.3 9.8l2.8-2m7.8 2l-2.8-2M3 3l18 18"/>`),
		g.Group(children),
	)
}