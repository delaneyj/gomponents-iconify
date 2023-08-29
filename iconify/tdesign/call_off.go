package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.414 3L11.833 13.58c.508.441 1.049.846 1.617 1.211l1.86-1.859L22 14.42V23h-1a19.91 19.91 0 0 1-10.85-3.197a20.085 20.085 0 0 1-2.567-1.971L3 22.414L1.586 21L21 1.586L22.414 3ZM9 16.415a18.069 18.069 0 0 0 2.237 1.71A17.892 17.892 0 0 0 20 20.972v-4.949l-4.053-.9l-2.174 2.175l-.663-.377A16.038 16.038 0 0 1 10.415 15L9 16.415ZM1 2h8.58l1.487 6.69l-1.865 1.866l.297.504l-1.723 1.015l-1.084-1.839l2.184-2.183L7.976 4H3.027a17.893 17.893 0 0 0 3.097 9.138l.564.825l-1.652 1.128l-.564-.825A19.911 19.911 0 0 1 1 3V2Z"/>`),
		g.Group(children),
	)
}