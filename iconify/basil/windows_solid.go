package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M13.027 10.507V5.122l7.453-1.235v6.62h-7.454Zm7.453 9.606L13.026 18.9v-5.405h7.454v6.618ZM9.633 10.505H3.565V6.622l6.068-1.005v4.888Zm0 7.907l-6.068-.989v-3.928h6.068v4.917Z"/>`),
		g.Group(children),
	)
}