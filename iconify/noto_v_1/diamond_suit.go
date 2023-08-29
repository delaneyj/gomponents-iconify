package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#db4437" d="M113.99 62.45L65.95.97c-.94-1.2-2.96-1.2-3.9 0L14.01 62.44c-.7.9-.7 2.16 0 3.05l48.04 61.48a2.478 2.478 0 0 0 3.9 0l48.03-61.48c.7-.89.7-2.15.01-3.04z"/>`),
		g.Group(children),
	)
}