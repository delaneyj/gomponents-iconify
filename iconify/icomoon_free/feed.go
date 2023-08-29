package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 8a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 8zm4.38-4.398a4.999 4.999 0 0 1 0 8.796c.689-1.096 1.12-2.66 1.12-4.398s-.431-3.302-1.12-4.398zM4.5 8c0 1.738.431 3.302 1.12 4.398a4.999 4.999 0 0 1 0-8.796C4.931 4.698 4.5 6.262 4.5 8zm-3 0c0 2.686.85 5.097 2.198 6.746C1.475 13.325 0 10.835 0 8s1.474-5.325 3.698-6.746C2.35 2.903 1.5 5.314 1.5 8zm10.802-6.746C14.525 2.675 16 5.165 16 8s-1.474 5.325-3.698 6.746C13.65 13.097 14.5 10.686 14.5 8s-.85-5.097-2.198-6.746z"/>`),
		g.Group(children),
	)
}