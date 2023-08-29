package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SsdRoundBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.5 13.588h13c1.311 0 2.454.764 3.053 1.893L19 5.118C18.5 3.529 17.605 3 16.5 3h-9c-1.105 0-2 .53-2.5 2.118L2.447 15.48c.6-1.13 1.742-1.893 3.053-1.893Z" opacity=".5"/><path fill-rule="evenodd" d="M18.5 13.588h-13c-1.311 0-2.454.763-3.053 1.893A3.86 3.86 0 0 0 2 17.294C2 19.34 3.567 21 5.5 21h13c1.933 0 3.5-1.66 3.5-3.706a3.86 3.86 0 0 0-.447-1.813c-.6-1.13-1.742-1.893-3.053-1.893ZM18 16.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1a.75.75 0 0 1 .75-.75Zm-1.75.75a.75.75 0 0 0-1.5 0v1a.75.75 0 0 0 1.5 0v-1ZM13 16.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1a.75.75 0 0 1 .75-.75Zm-1.75.75a.75.75 0 0 0-1.5 0v1a.75.75 0 0 0 1.5 0v-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}