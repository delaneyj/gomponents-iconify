package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoodreadsLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 104v8a32 32 0 0 1-64 0v-8a32 32 0 0 1 64 0Zm72 24A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-56-64a8 8 0 0 0-16 0v4.26A48 48 0 0 0 80 104v8a48 48 0 0 0 80 35.74V160a32 32 0 0 1-57.6 19.2a8 8 0 1 0-12.8 9.6A48 48 0 0 0 176 160Z"/>`),
		g.Group(children),
	)
}