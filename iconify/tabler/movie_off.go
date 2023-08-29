package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h10a2 2 0 0 1 2 2v10m-.592 3.42c-.362.359-.859.58-1.408.58H6a2 2 0 0 1-2-2V6c0-.539.213-1.028.56-1.388M8 8v12m8-16v8m0 4v4M4 8h4m-4 8h4m-4-4h8m4 0h4m-4-4h4M3 3l18 18"/>`),
		g.Group(children),
	)
}