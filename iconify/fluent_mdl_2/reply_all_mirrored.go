package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAllMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M742 640q-128 0-245 48T289 827q-91 91-139 208t-48 245q0 133 50 249t137 204t203 137t250 50v-128q-106 0-199-40t-162-110t-110-163t-41-199q0-106 40-199t110-162t163-110t199-41h549l-402 403l90 90l557-557l-557-557l-90 90l402 403H742zm1126 64l-467 467l90 90l557-557l-557-557l-90 90l467 467z"/>`),
		g.Group(children),
	)
}