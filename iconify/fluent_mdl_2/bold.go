package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1255 901q92 0 173 35t142 95t96 142t35 173q0 92-35 173t-95 142t-142 96t-174 35H512V128h743q80 0 150 30t122 83t83 123t31 150q0 80-30 150t-82 123t-123 83t-151 31zm-208-535H896v475h151q49 0 92-19t76-51t51-75t19-93q0-49-19-92t-51-75t-75-51t-93-19zm59 1188q49 0 92-18t76-51t51-76t19-92q0-49-18-92t-51-76t-76-51t-93-19H896v475h210z"/>`),
		g.Group(children),
	)
}