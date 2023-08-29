package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tangerine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 2H7a5 5 0 0 0 4.625 4.986C7.144 8.886 4 13.326 4 18.5C4 25.404 9.596 31 16.5 31S29 25.404 29 18.5c0-6.77-5.38-12.282-12.099-12.494A5.002 5.002 0 0 0 12 2ZM6 18.5C6 12.701 10.701 8 16.5 8S27 12.701 27 18.5S22.299 29 16.5 29S6 24.299 6 18.5Z"/>`),
		g.Group(children),
	)
}