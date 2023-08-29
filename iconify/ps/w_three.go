package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m172 20l61 209l62-209h149v22l-54 108q37 12 53 38q19 30 19 66q0 48-23 78q-26 32-64 32q-27 0-51-19q-22-20-32-51l35-15q8 22 20 31q11 11 29 11q19 0 30-20q12-19 12-47q0-38-13-54q-19-26-48-26h-29v-8l72-107h-76l-86 287h-5l-63-213l-63 213h-4L2 20h44l61 209l41-140l-20-69h44z"/>`),
		g.Group(children),
	)
}