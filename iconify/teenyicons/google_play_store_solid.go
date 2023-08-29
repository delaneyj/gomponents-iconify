package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayStoreSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.251.066a.5.5 0 0 1 .5.002l7.843 4.575l-2.427 2.184L1 1.277V.5a.5.5 0 0 1 .251-.434ZM1 2.623v9.754L6.42 7.5L1 2.623Zm0 11.1v.777a.5.5 0 0 0 .752.432l7.842-4.575l-2.427-2.184L1 13.723Zm9.501-3.895l3.25-1.896a.5.5 0 0 0 0-.864l-3.25-1.896L7.914 7.5l2.587 2.328Z"/>`),
		g.Group(children),
	)
}