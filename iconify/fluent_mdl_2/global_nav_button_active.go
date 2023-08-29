package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobalNavButtonActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1728 1024q-66 0-124-25t-102-68t-69-102t-25-125q0-66 25-124t68-102t102-69t125-25q66 0 124 25t102 68t69 102t25 125q0 66-25 124t-68 102t-102 69t-125 25zm-474-512q-12 31-19 63t-13 65H0V512h1254zm78 512q65 80 153 128H0v-128h1332zM0 1664v-128h2048v128H0z"/>`),
		g.Group(children),
	)
}