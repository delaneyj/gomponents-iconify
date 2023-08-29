package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2.75a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5ZM11.25 4.5a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Zm-.747 4.707a2.25 2.25 0 0 1 3.136 2.068v2.62c0 .47-.12.933-.35 1.343l-2.712 4.848A2.75 2.75 0 0 1 9.1 21.334l-3.849 1.372a.75.75 0 1 1-.504-1.412l3.849-1.373a1.25 1.25 0 0 0 .67-.567l2.713-4.848a1.25 1.25 0 0 0 .16-.61v-2.621a.75.75 0 0 0-1.046-.69l-2.456 1.053a.75.75 0 0 0-.228 1.226l.614.599a.75.75 0 0 1-1.047 1.074l-.614-.599a2.25 2.25 0 0 1 .684-3.679l2.456-1.052Zm6.056 4.846a2.75 2.75 0 0 1 3.191 2.715V22a.75.75 0 0 1-1.5 0v-5.232c0-.77-.69-1.357-1.45-1.234l-1.013.165a.75.75 0 1 1-.24-1.481l1.012-.165Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}