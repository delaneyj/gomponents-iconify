package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitbucketSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5 0a.5.5 0 0 0-.495.57l2 14A.5.5 0 0 0 2.5 15h10a.5.5 0 0 0 .495-.43L14.219 6H9.373l-.333 4H5.96l-.417-5h8.82l.632-4.43A.5.5 0 0 0 14.5 0H.5Z"/>`),
		g.Group(children),
	)
}