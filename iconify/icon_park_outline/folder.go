package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M7 6a2 2 0 0 1 2-2h30a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V6Zm9 23h4m-4 6h10"/><path d="M8 5s3.765 13 16 13S40 5 40 5"/><circle cx="24" cy="18" r="4"/><path d="M15 4H9a2 2 0 0 0-2 2v6m26-8h6a2 2 0 0 1 2 2v6"/></g>`),
		g.Group(children),
	)
}