package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatUnderlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 299q-53 0-90.5-37.5T21 171V0h54v171q0 31 21.5 52.5T149 245t53-21.5t22-52.5V0h53v171q0 53-37.5 90.5T149 299zM0 341h299v43H0v-43z"/>`),
		g.Group(children),
	)
}