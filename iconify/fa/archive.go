package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 704q0-26-19-45t-45-19H704q-26 0-45 19t-19 45t19 45t45 19h256q26 0 45-19t19-45zm576-192v960q0 26-19 45t-45 19H128q-26 0-45-19t-19-45V512q0-26 19-45t45-19h1408q26 0 45 19t19 45zm64-448v256q0 26-19 45t-45 19H64q-26 0-45-19T0 320V64q0-26 19-45T64 0h1536q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}