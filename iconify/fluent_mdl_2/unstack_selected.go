package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnstackSelected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 640v1152H512v-256H256v-256H0V128h1313q-69 20-133 52t-123 76H128v896h1280V640h64l64 64v576H384v128h1280V768l128-128v896H640v128h1280V640h128zM896 896q0-88 23-170t64-153t100-129t130-100t153-65t170-23h165l-146-147l90-90l301 301l-301 301l-90-90l146-147h-165q-106 0-199 40t-162 110t-110 163t-41 199H896z"/>`),
		g.Group(children),
	)
}