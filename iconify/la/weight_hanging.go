package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeightHanging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5c-1.645 0-3 1.355-3 3c0 .352.074.684.188 1h-5l-.157.813l-3 15l-.031.093V27h22v-2.094l-.031-.093l-3-15L23.812 9h-5A2.95 2.95 0 0 0 19 8c0-1.645-1.355-3-3-3zm0 2c.563 0 1 .438 1 1c0 .563-.438 1-1 1c-.563 0-1-.438-1-1c0-.563.438-1 1-1zm-6.188 4h12.376L25 25H7z"/>`),
		g.Group(children),
	)
}