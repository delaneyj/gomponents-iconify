package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m880 1329l80 237q-23 35-111 66t-177 32q-104 2-190.5-26T339 1564t-95-106t-55.5-120t-16.5-118V676H4V461q72-26 129-69.5t91-90t58-102t34-99T331 12q1-5 4.5-8.5T343 0h244v424h333v252H586v518q0 30 6.5 56t22.5 52.5t49.5 41.5t81.5 14q78-2 134-29z"/>`),
		g.Group(children),
	)
}