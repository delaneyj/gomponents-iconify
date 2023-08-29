package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenZipper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.026 21.105V6.5h1.948v14.605m-2.921.027v1.947l.974.974h1.948l.974-.974v-1.947"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m24 24.053l-.974.974v1.948l.974.974l.974-.974v-1.948Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.053 21.132h3.895m-2.922-2.948h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948m-1.948-2.921h1.948M18 34.7h4L18 40h4m-9.5 0l4.3-8h-5.3"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M37.5 4.5h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}