package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 256q40 0 75 15t61 41t41 61t15 75v768h-128V448q0-26-19-45t-45-19H384v1280h832v128H0V256h1856zM256 384H128v1280h128V384zm1280 640H768V640h768v384zm-128-256H896v128h512V768zm586 685l-226 227l226 226l-90 90l-226-226l-225 224l-90-90l224-224l-224-225l90-90l225 224l226-226l90 90z"/>`),
		g.Group(children),
	)
}