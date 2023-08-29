package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v768h-128V648l-896 447l-896-447v888h1280v128H0V384zm143 128l881 441l881-441H143zm1649 1536v-768h128v768h-128zm-256 0v-768h128v768h-128z"/>`),
		g.Group(children),
	)
}