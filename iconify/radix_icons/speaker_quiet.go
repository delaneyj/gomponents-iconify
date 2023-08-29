package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerQuiet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a.5.5 0 0 0-.8-.4L3.333 4H1.5A1.5 1.5 0 0 0 0 5.5v4A1.5 1.5 0 0 0 1.5 11h1.833L7.2 13.9a.5.5 0 0 0 .8-.4v-12ZM3.8 4.9L7 2.5v10l-3.2-2.4a.5.5 0 0 0-.3-.1h-2a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 0 .3-.1Zm6.283.156a.4.4 0 1 0-.666.443a3.623 3.623 0 0 1 0 4.002a.4.4 0 1 0 .666.443a4.423 4.423 0 0 0 0-4.888Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}