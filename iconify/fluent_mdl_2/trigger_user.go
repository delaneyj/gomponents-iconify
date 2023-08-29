package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriggerUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1883 768L603 2048H313l384-768H248L888 0h719l-384 768h660zm-310 128h-557l384-768H967L455 1152h449l-384 768h29L1573 896zm275 817q46 25 83 61t63 79t40 93t14 102h-128q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100h-128q0-52 14-101t40-93t63-80t83-61q-34-35-53-81t-19-96q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100q0 50-19 96t-53 81zm-184-49q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10z"/>`),
		g.Group(children),
	)
}