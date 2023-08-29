package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialLinkedin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19H5V9h3v10zm11 0h-3v-5.342c0-1.392-.496-2.085-1.479-2.085c-.779 0-1.273.388-1.521 1.165V19h-3s.04-9 0-10h2.368l.183 2h.062c.615-1 1.598-1.678 2.946-1.678c1.025 0 1.854.285 2.487 1.001c.637.717.954 1.679.954 3.03V19z"/><ellipse cx="6.5" cy="6.5" fill="currentColor" rx="1.55" ry="1.5"/>`),
		g.Group(children),
	)
}