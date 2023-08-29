package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.532 33.313a6.174 6.174 0 0 1-11.493 3.142l-.124-.217L24 24.23l-5.266-9.145l-.174-.291l-.043-.081a6.158 6.158 0 0 1-.694-2.856c0-1.047.26-2.039.725-2.906l.167-.291A6.168 6.168 0 0 1 24 5.68c2.503 0 4.09.948 5.626 3.624v.006l12.014 20.8l.068.111a6.13 6.13 0 0 1 .824 3.092Zm-24.709 0c0 1.066-.273 2.07-.75 2.95l-.112.198a6.174 6.174 0 0 1-5.316 3.03a6.174 6.174 0 0 1-6.177-6.178c0-1.04.254-2.02.713-2.88l.235-.41l.006-.012a6.178 6.178 0 0 1 11.4 3.302ZM24 24.23l-6.927 12.032M6.422 30.01L18.548 8.95"/>`),
		g.Group(children),
	)
}