package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func R(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1023 1664H896V896h253q49 0 94 12t80 37t55 64t21 94q0 83-46 136t-124 76q39 22 70 58t54 75l136 216h-148l-123-206q-13-22-27-41t-30-33t-38-23t-49-9h-51v312zm0-415h106q60 0 98-35t38-97q0-62-39-90t-96-28h-107v250zm897-1121v1792H128V128h1792zm-128 128H256v1536h1536V256z"/>`),
		g.Group(children),
	)
}