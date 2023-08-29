package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M56.089 1.104s-48.341-.015-48.326 0c-4.455 0-6.96 2.303-6.96 7.009V56.63c0 4.402 2.252 6.707 6.704 6.707h48.722c4.452 0 6.706-2.19 6.706-6.707V8.113c.001-4.592-2.254-7.009-6.846-7.009zM52.34 39.112H39.263v13.077H24.214V39.112H11.137V24.064h13.077V10.987h15.049v13.077H52.34v15.048z"/>`),
		g.Group(children),
	)
}