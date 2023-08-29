package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 324L512 196v-4q0-40 15-75t41-61t61-41t75-15h512q40 0 75 15t61 41t41 61t15 75v900l-128-128V192q0-27-18-45t-46-19H704q-27 0-45 18t-19 46v132zm1024 700v323l-128-127v-196h128zm365 915l-90 90l-377-377q-59 67-139 103t-169 37h-230v128h256v128H640v-128h256v-128H666q-85 0-159-32t-130-88t-88-130t-33-160v-358h128v358q0 58 22 109t61 90t89 60t110 23h588q63 0 119-27t98-75l-102-103q-27 36-67 56t-86 21H704q-40 0-75-15t-61-41t-41-61t-15-75V603L19 109l90-90l1920 1920zm-813-531q20 0 36-11t24-30L640 731v613q0 27 18 45t46 19h512z"/>`),
		g.Group(children),
	)
}