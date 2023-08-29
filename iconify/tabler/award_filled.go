package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m19.496 13.983l1.966 3.406a1.001 1.001 0 0 1-.705 1.488l-.113.011l-.112-.001l-2.933-.19l-1.303 2.636a1.001 1.001 0 0 1-1.608.26l-.082-.094l-.072-.11l-1.968-3.407a8.994 8.994 0 0 0 6.93-3.999zm-8.066 3.999L9.464 21.39a1.001 1.001 0 0 1-1.622.157l-.076-.1l-.064-.114l-1.304-2.635l-2.931.19a1.001 1.001 0 0 1-1.022-1.29l.04-.107l.05-.1l1.968-3.409a8.994 8.994 0 0 0 6.927 4.001zM12 2l.24.004A7 7 0 0 1 19 9l-.003.193l-.007.192l-.018.245l-.026.242l-.024.178a6.985 6.985 0 0 1-.317 1.268l-.116.308l-.153.348a7.001 7.001 0 0 1-12.688-.028l-.13-.297l-.052-.133l-.08-.217l-.095-.294a6.96 6.96 0 0 1-.093-.344l-.06-.271l-.049-.271l-.02-.139l-.039-.323l-.024-.365L5 9a7 7 0 0 1 6.76-6.996L12 2z"/></g>`),
		g.Group(children),
	)
}