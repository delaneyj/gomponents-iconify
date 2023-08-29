package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationVirtual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 19h4v4h-4zm-6 0h4v4h-4zm6 6h4v4h-4zm-6 0h4v4h-4z"/><path fill="currentColor" d="M24.5 25H24v-2h.5a5.496 5.496 0 0 0 .377-10.98l-.837-.056l-.09-.834a7.998 7.998 0 0 0-15.9 0l-.09.834l-.837.057A5.496 5.496 0 0 0 7.5 23H8v2h-.5a7.496 7.496 0 0 1-1.322-14.876a10 10 0 0 1 19.644 0A7.496 7.496 0 0 1 24.5 25Z"/>`),
		g.Group(children),
	)
}