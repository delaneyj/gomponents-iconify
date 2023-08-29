package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.643 16.757L10.499 15.5h-4v-7h2L11.3 3c.58 0 1.075.205 1.485.615c.41.41.615.905.615 1.485l-.9 2.4l4.031 1.344a2 2 0 0 1 1.309 2.38l-.069.22l-1.553 4.142a2 2 0 0 1-2.575 1.17z"/><path d="M3.5 7.5h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}