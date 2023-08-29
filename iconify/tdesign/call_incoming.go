package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallIncoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.032 3.378L18.738 6.67l2.12.046l-.044 2l-5.418-.117l-.117-5.419l2-.043l.045 2.119l3.293-3.293l1.415 1.414ZM1 2h8.58l1.487 6.69l-1.86 1.86a14.081 14.081 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.911 19.911 0 0 1-10.85-3.196a20.102 20.102 0 0 1-5.954-5.954A19.911 19.911 0 0 1 1 3V2Zm2.027 2a17.893 17.893 0 0 0 2.849 8.764a18.102 18.102 0 0 0 5.36 5.36A17.893 17.893 0 0 0 20 20.974v-4.948l-4.053-.901l-2.174 2.175l-.663-.377a16.072 16.072 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4H3.027Z"/>`),
		g.Group(children),
	)
}