package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriggerApproval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1883 768L603 2048H313l384-768H248L888 0h719l-384 768h660zm-310 128h-557l384-768H967L455 1152h449l-384 768h29L1573 896zm456 557l-557 558l-269-270l90-90l179 178l467-466l90 90z"/>`),
		g.Group(children),
	)
}