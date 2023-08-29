package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1728 640q-66 0-124-25t-102-68t-69-102t-25-125q0-66 25-124t68-102t102-69t125-25q66 0 124 25t102 68t69 102t25 125q0 66-25 124t-68 102t-102 69t-125 25zm-576 256V640h218q18 20 38 38v218q0 27-10 50t-27 40t-41 28t-50 10V896h-128zm-384 128V896H640V640h256v256q0 27-10 50t-27 40t-41 28t-50 10zm1152 384V760q35-15 67-36t61-46v858H731l-475 475v-475H0V128h1288q-27 61-36 128H128v1152h256v293l293-293h1243z"/>`),
		g.Group(children),
	)
}