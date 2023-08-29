package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonStarsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 96a12 12 0 0 1-12 12h-12v12a12 12 0 0 1-24 0v-12h-12a12 12 0 0 1 0-24h12V72a12 12 0 0 1 24 0v12h12a12 12 0 0 1 12 12ZM144 60h4v4a12 12 0 0 0 24 0v-4h4a12 12 0 0 0 0-24h-4v-4a12 12 0 0 0-24 0v4h-4a12 12 0 0 0 0 24Zm75.81 90.38A12 12 0 0 1 222 162.3A100 100 0 1 1 93.7 34a12 12 0 0 1 15.89 13.6A85.12 85.12 0 0 0 108 64a84.09 84.09 0 0 0 84 84a85.22 85.22 0 0 0 16.37-1.59a12 12 0 0 1 11.44 3.97ZM190 172A108.13 108.13 0 0 1 84 66a76 76 0 1 0 106 106Z"/>`),
		g.Group(children),
	)
}