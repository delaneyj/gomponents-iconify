package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodcastOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.264 9.557a6.75 6.75 0 0 1 13.472 0A2.75 2.75 0 0 1 22.75 12v2a2.75 2.75 0 0 1-2.904 2.746v.004a2.596 2.596 0 0 1-2.596-2.596V10a5.25 5.25 0 0 0-10.5 0v4.154a2.596 2.596 0 0 1-2.596 2.596v-.004A2.75 2.75 0 0 1 1.25 14v-2a2.75 2.75 0 0 1 4.014-2.443ZM5.25 12a1.25 1.25 0 0 0-2.5.001v2a1.25 1.25 0 0 0 2.5 0s0 .001 0 0v-2s0-.001 0 0Zm3-1.999a3.75 3.75 0 1 1 7.5 0v3a3.751 3.751 0 0 1-3 3.675V19a.75.75 0 0 1-1.5 0v-2.325a3.751 3.751 0 0 1-3-3.675v-3Zm1.628-.75H10a.75.75 0 0 1 0 1.5h-.25v1.5H10a.75.75 0 0 1 0 1.5h-.122a2.25 2.25 0 0 0 4.244 0H13a.75.75 0 0 1 0-1.5h1.25v-1.5H13a.75.75 0 0 1 0-1.5h1.122a2.251 2.251 0 0 0-4.244 0ZM20 10.75c-.69 0-1.25.56-1.25 1.25v2a1.25 1.25 0 1 0 2.5 0v-2c0-.69-.56-1.25-1.25-1.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}