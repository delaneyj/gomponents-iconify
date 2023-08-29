package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M403 3q67 0 127 24t106 68t74 102t31 126q3 72-22 136t-71 112t-108 76t-135 28q-58 0-110-19t-95-51q-4-3-4-9t3-9l50-49q7-7 15-2q30 22 66 34t75 12q54 0 101-22t81-59t50-86t10-104q-5-41-23-78t-46-65t-64-47t-78-24q-54-6-102 9t-86 48t-60 76t-25 97h69L116 466L0 327h70q2-67 29-126t72-103t105-70T403 3z"/>`),
		g.Group(children),
	)
}