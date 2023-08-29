package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TclAr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-9.44 11.904h8.74m-4.37 13.192V17.404m-15.5 0l8.74 13.192m0-13.192l-8.74 13.192"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.8 30.596V17.404l8.74 13.192V17.404"/>`),
		g.Group(children),
	)
}