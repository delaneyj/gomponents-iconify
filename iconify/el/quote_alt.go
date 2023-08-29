package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zm-45.923 282.422l25.781 118.727c-75.541 16.721-145.005 38.468-139.38 122.826h88.77v315.968H261.841V544.629c.095-234.691 172.401-253.786 292.236-262.207zm358.228 0l25.854 118.727c-75.541 16.721-145.005 38.468-139.38 122.826h88.77v315.968H620.142V544.629c.094-234.691 172.328-253.786 292.163-262.207z"/>`),
		g.Group(children),
	)
}