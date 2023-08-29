package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GIF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 128q39 0 74 15t61 41t42 62t15 74v1280q0 39-15 74t-41 61t-62 42t-74 15H192q-39 0-74-15t-61-41t-42-62t-15-74V320q0-39 15-74t41-61t62-42t74-15h1664zm64 1472V320q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v1280q0 26 19 45t45 19h1664q26 0 45-19t19-45zm-626-326V646h338v89h-234v188h216v88h-216v263h-104zm-656-254v-87h236v298q-51 27-106 40t-112 13q-71 0-128-23t-99-64t-63-100t-22-129q0-73 25-134t69-105t106-68t134-25q45 0 89 6t87 24v108q-38-26-82-37t-90-12q-52 0-94 18t-72 50t-46 75t-16 95q0 48 12 90t38 74t64 50t90 18q29 0 57-5t55-20v-150H638zm396 260V653h108v627h-108z"/>`),
		g.Group(children),
	)
}