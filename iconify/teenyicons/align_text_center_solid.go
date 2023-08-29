package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextCenterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 3h15v1H0V3Zm5 4h5v1H5V7Zm-2 4h9v1H3v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}