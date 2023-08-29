package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAllAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 896v128H859l402 403l-90 90l-557-557l557-557l90 90l-402 403h1061zM749 493L282 960l467 467l-90 90l-557-557l557-557l90 90z"/>`),
		g.Group(children),
	)
}