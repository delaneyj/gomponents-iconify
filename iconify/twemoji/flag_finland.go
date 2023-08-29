package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFinland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EDECEC" d="M32 5H18v10h18V9a4 4 0 0 0-4-4z"/><path fill="#EEE" d="M11 5H4a4 4 0 0 0-4 4v6h11V5z"/><path fill="#EDECEC" d="M32 31H18V21h18v6a4 4 0 0 1-4 4zm-21 0H4a4 4 0 0 1-4-4v-6h11v10z"/><path fill="#003580" d="M18 5h-7v10H0v6h11v10h7V21h18v-6H18z"/>`),
		g.Group(children),
	)
}