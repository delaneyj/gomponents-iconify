package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v85L1024 981L0 469v-85h2048zm-1024 811l1024-512v981H0V683l1024 512z"/>`),
		g.Group(children),
	)
}