package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThisPC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h1920v1152h-896v128h256v128H640v-128h256v-128H0V0zm1792 1024V128H128v896h1664zM228 1536h1464l228 334v50H0v-50l228-334zm-20 256h1504l-88-128H296l-88 128z"/>`),
		g.Group(children),
	)
}