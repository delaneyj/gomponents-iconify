package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirebaseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.33.03a.5.5 0 0 1 .524.116l2.078 2.08a.505.505 0 0 0-.032.056L2.175 9.988L3.33.03ZM2.262 11.94l4.98 2.989a.5.5 0 0 0 .444.035l5-2a.5.5 0 0 0 .31-.52l-1-9a.5.5 0 0 0-.828-.318L9.513 4.597L2.262 11.94Zm6.676-8.184L7.916 2.223a.5.5 0 0 0-.853.034l-.31.558l-2.986 6.177l5.171-5.236Z"/>`),
		g.Group(children),
	)
}