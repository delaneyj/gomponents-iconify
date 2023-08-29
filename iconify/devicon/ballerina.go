package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ballerina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#46C0BC" d="M29 35.925V0h30.082v47.43L29 35.925Zm0 26.995l20.59-7.873L29 47.172V62.92Zm0 11.247V128h15.458l14.624-47.936v-17.4L29 74.166Zm70.587-38.242V0H69.506v47.43l30.081-11.505Zm0 11.248l-20.589 7.873l20.59 7.874V47.172Zm-30.081 15.49v17.4L84.129 128h15.458V74.167L69.506 62.663Z"/>`),
		g.Group(children),
	)
}