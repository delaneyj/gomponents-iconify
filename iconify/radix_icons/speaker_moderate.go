package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerModerate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a.5.5 0 0 0-.8-.4L3.333 4H1.5A1.5 1.5 0 0 0 0 5.5v4A1.5 1.5 0 0 0 1.5 11h1.833L7.2 13.9a.5.5 0 0 0 .8-.4v-12ZM3.8 4.9L7 2.5v10l-3.2-2.4a.5.5 0 0 0-.3-.1h-2a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 0 .3-.1Zm7.033-.94a.4.4 0 0 0-.666.443a5.607 5.607 0 0 1 0 6.194a.4.4 0 1 0 .666.444a6.407 6.407 0 0 0 0-7.082Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}