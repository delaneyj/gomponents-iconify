package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilePhoneWithArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13 2a3 3 0 0 0-3 3v7.79l1 .91V6h17v20H11v-7.76l-1 .91V27a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H13Z"/><path d="M7.04 11.45a.518.518 0 0 0-.87.39v1.81c0 .29-.23.52-.52.52H3.76a.76.76 0 0 0-.76.76v2.08c0 .42.34.76.76.76h1.89c.29 0 .52.23.52.52v1.81c0 .45.54.69.87.39l4.54-4.13c.23-.21.23-.57 0-.78l-4.54-4.13Zm5.46 3.678a1.53 1.53 0 0 1 0 1.684V17a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v.128ZM18.5 9a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2Zm0 10a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2Zm-5-10a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2Zm-1 11a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2Zm10-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2Zm1-6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2Zm-6 6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2Zm5 5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2Z"/></g>`),
		g.Group(children),
	)
}