package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1468 476q14 14 28 36h-472V40q22 14 36 28zM992 640h544v1056q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h800v544q0 40 28 68t68 28zm160 736v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23zm0-256v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23zm0-256v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23z"/>`),
		g.Group(children),
	)
}