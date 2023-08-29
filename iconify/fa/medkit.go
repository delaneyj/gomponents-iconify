package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 992V800q0-14-9-23t-23-9h-224V544q0-14-9-23t-23-9H800q-14 0-23 9t-9 23v224H544q-14 0-23 9t-9 23v192q0 14 9 23t23 9h224v224q0 14 9 23t23 9h192q14 0 23-9t9-23v-224h224q14 0 23-9t9-23zM640 256h512V128H640v128zm-384 0v1280h-32q-92 0-158-66T0 1312V480q0-92 66-158t158-66h32zm1184 0v1280H352V256h160V96q0-40 28-68t68-28h576q40 0 68 28t28 68v160h160zm352 224v832q0 92-66 158t-158 66h-32V256h32q92 0 158 66t66 158z"/>`),
		g.Group(children),
	)
}