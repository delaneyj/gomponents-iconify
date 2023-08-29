package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.27 3.388c-.991-1.85-3.663-1.85-4.643 0l-6.273 11.8c-3.842 6.58.95 14.81 8.614 14.81h.21c7.554 0 12.266-8.1 8.494-14.58L18.27 3.388Z"/>`),
		g.Group(children),
	)
}