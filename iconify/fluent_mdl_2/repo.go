package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 768h-128v1152H128V768H0V128h2048v640zm-256 0h-512v768l-256-128l-256 128V768H256v1024h1536V768zm-896 561q32-16 64-31t64-33q32 17 64 32t64 32V768H896v561zM1920 256H128v384h1792V256z"/>`),
		g.Group(children),
	)
}