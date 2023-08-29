package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v9h-2V5h-2v2H7V5H5v16h7v2H3V3h4V1Zm2 4h6V3H9v2Zm11 8.75v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4.008 4.008 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 20 22.874v1.376h-2v-1.376a3.996 3.996 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4.004 4.004 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.689A3.996 3.996 0 0 1 18 15.126V13.75h2Zm-2.751 4.283a1.991 1.991 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a1.999 1.999 0 0 0 3.43 0l.036-.063A1.99 1.99 0 0 0 21 19c0-.35-.09-.68-.249-.967l-.036-.063a1.999 1.999 0 0 0-3.43 0l-.036.063Z"/>`),
		g.Group(children),
	)
}