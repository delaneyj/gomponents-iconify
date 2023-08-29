package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingNotAllowedSpaceMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm5.25 6.75a5.25 5.25 0 1 0-10.5 0v2.25H9l.75 7.5h4.5l.75-7.5h2.25V13.5ZM1 18.991l4 4m0-4l-4 4m18-4l4 4m0-4l-4 4"/>`),
		g.Group(children),
	)
}