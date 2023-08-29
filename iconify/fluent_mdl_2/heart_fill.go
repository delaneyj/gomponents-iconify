package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1889 287q53 53 88 116t53 131t18 138t-17 138t-53 131t-89 116l-865 864l-865-864q-53-53-88-116T18 811T0 673t17-139t53-131t89-116q78-77 177-118t208-41q109 0 208 41t177 118l95 96l95-96q78-77 177-118t208-41q109 0 208 41t177 118z"/>`),
		g.Group(children),
	)
}