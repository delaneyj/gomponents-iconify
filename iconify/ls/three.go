package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Three(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M105 192H33C35 87 122 0 228 0s193 87 193 195c0 63-31 119-78 154c67 40 112 113 112 196c0 126-102 227-227 227C105 772 5 676 0 555h72c5 81 73 147 156 147c86 0 155-71 155-157s-69-156-155-156v-71c68 0 122-55 122-123S296 72 228 72s-122 53-123 120z"/>`),
		g.Group(children),
	)
}