package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertBridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="23.766" height="26.803" x="12.117" y="10.598" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.726"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.398 10.598l-.758-5.356a.863.863 0 0 0-.855-.742h-13.57a.863.863 0 0 0-.854.742l-.759 5.357m0 26.802l.758 5.357a.863.863 0 0 0 .855.742h13.57a.863.863 0 0 0 .854-.742l.759-5.357m-1.726-21.799H17.328a1.726 1.726 0 0 0-1.726 1.726v14.785l3.388-3.388h11.682a1.726 1.726 0 0 0 1.726-1.727v-9.67a1.726 1.726 0 0 0-1.726-1.726Z"/><circle cx="20.515" cy="22.191" r=".75" fill="currentColor"/><circle cx="24" cy="22.191" r=".75" fill="currentColor"/><circle cx="27.485" cy="22.191" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}