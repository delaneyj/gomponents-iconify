package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WriteOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 242.9h170.7V72.3H0v170.6zM213.3 93.6v42.7H512V93.6H213.3zm0 128H512v-42.7H213.3v42.7zm0 128H512v-42.7H213.3v42.7zm0 85.3H512v-42.7H213.3v42.7zM0 456.3h170.7V285.6H0v170.7z"/>`),
		g.Group(children),
	)
}