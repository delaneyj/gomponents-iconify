package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hyper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M289.381 382H512v42H289.381v-42zm-68.3-108.52L25.106 444.233l60.997-149.45L0 238.503L195.992 67.768l-66.34 158.968l91.429 46.745z"/>`),
		g.Group(children),
	)
}