package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M38 14h590c20 0 38 18 38 38v590c0 20-18 38-38 38H38c-20 0-38-18-38-38V52c0-20 18-38 38-38zm435 374h92l4-88h-96v-65c0-25 5-39 37-39h56l2-82s-25-4-61-4c-88 0-127 55-127 114v76h-65v88h65v244h93V388z"/>`),
		g.Group(children),
	)
}