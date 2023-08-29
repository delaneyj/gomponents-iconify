package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M768 832v448q0 26-19 45t-45 19t-45-19l-144-144l-332 332q-10 10-23 10t-23-10L23 1399q-10-10-10-23t10-23l332-332l-144-144q-19-19-19-45t19-45t45-19h448q26 0 45 19t19 45zm755-672q0 13-10 23l-332 332l144 144q19 19 19 45t-19 45t-45 19H832q-26 0-45-19t-19-45V256q0-26 19-45t45-19t45 19l144 144l332-332q10-10 23-10t23 10l114 114q10 10 10 23z"/>`),
		g.Group(children),
	)
}