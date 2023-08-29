package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CondicionerTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8 5c-2.828 0-4.243 0-5.121.879C2 6.757 2 8.172 2 11c0 3.771 0 5.657 1.172 6.828C4.343 19 6.229 19 10 19h4c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-2.828 0-4.243-.879-5.121C20.243 5 18.828 5 16 5h-4"/><path d="M18 18.5c0-1.404 0-2.107-.337-2.611a1.999 1.999 0 0 0-.552-.552C16.607 15 15.904 15 14.5 15h-5c-1.404 0-2.107 0-2.611.337a2 2 0 0 0-.552.552C6 16.393 6 17.096 6 18.5m0-7h2m10 0h-6M6 9h6m6 0h-2"/></g>`),
		g.Group(children),
	)
}