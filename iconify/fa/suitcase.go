package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 256h512V128H640v128zm-352 0v1280h-64q-92 0-158-66T0 1312V480q0-92 66-158t158-66h64zm1120 0v1280H384V256h128V96q0-40 28-68t68-28h576q40 0 68 28t28 68v160h128zm384 224v832q0 92-66 158t-158 66h-64V256h64q92 0 158 66t66 158z"/>`),
		g.Group(children),
	)
}