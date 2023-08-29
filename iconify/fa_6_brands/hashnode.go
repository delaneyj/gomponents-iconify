package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashnode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M35.19 171.1c-46.91 46-46.91 122.9 0 169.8L171.1 476.8c46 46.9 122.9 46.9 169.8 0l135.9-135.9c46.9-46.9 46.9-123.8 0-169.8L340.9 35.19c-46.9-46.91-123.8-46.91-169.8 0L35.19 171.1zM315.5 315.5c-32.9 32.8-86.1 32.8-118.9 0c-32.9-32.9-32.9-86.1 0-118.9c32.8-32.9 86-32.9 118.9 0c32.8 32.8 32.8 86 0 118.9z"/>`),
		g.Group(children),
	)
}