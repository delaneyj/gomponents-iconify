package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerraformIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#4040B2" d="M176.485 188.994L256 143.127V51.249l-79.515 45.963z"/><path fill="#5C4EE5" d="m88.243 51.249l79.515 45.963v91.782L88.243 143.08M0 91.83l79.515 45.916v-91.83L0 0m88.243 244.994l79.515 45.915v-91.83l-79.515-45.915"/>`),
		g.Group(children),
	)
}