package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M870 640q-128 0-245 48T417 827q-91 91-139 208t-48 245q0 133 50 249t137 204t203 137t250 50v-128q-106 0-199-40t-162-110t-110-163t-41-199q0-106 40-199t110-162t163-110t199-41h677l-402 403l90 90l557-557l-557-557l-90 90l402 403H870z"/>`),
		g.Group(children),
	)
}