package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletSelected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1408h256v128H896v-128zm992-1152q33 0 62 12t51 35t34 51t13 62v843l-128-123V416q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v1216q0 14 9 23t23 9h800l128 128H160q-33 0-62-12t-51-35t-34-51t-13-62V416q0-33 12-62t35-51t51-34t62-13h1728zm51 1091l90 91l-557 557l-269-269l90-91l179 179l467-467z"/>`),
		g.Group(children),
	)
}