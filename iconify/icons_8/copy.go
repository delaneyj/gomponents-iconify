package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 4v20h7v-2H6V6h12v1h2V4H4zm8 4v20h16V8H12zm2 2h12v16H14V10z"/>`),
		g.Group(children),
	)
}