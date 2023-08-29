package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roadtunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 640q-13 0-22.5-9.5T704 608V384q0-87-43-160.5T544.5 107T384 64t-160.5 43T107 223.5T64 384v224q0 13-9.5 22.5T32 640t-22.5-9.5T0 608V384q0-104 51.5-192.5t140-140T384 0t192.5 51.5t140 140T768 384v224q0 13-9.5 22.5T736 640zm-576 0q-13 0-22.5-9.5T128 608V384q0-106 75-181t181-75t181 75t75 181v224q0 13-9.5 22.5T608 640h-32l96 384H96l96-384h-32zm266 128h-84l-7 64h98zm22 192l-8-64H328l-8 64h128zm-99-256h70l-3-64h-64z"/>`),
		g.Group(children),
	)
}