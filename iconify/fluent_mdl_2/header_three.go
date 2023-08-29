package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeaderThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 384h128v1408H768v-640H128v640H0V384h128v640h640V384zm851 674q67 7 125 32t102 67t69 99t25 127q0 103-39 179t-106 128t-153 76t-181 26q-80 0-160-15t-149-59v-167q138 108 315 108q61 0 118-15t100-48t70-81t26-118q0-69-22-116t-61-76t-88-45t-105-22t-111-7t-106-1V998h98q51 0 101-7t94-20t77-43t53-71t20-109q0-62-18-106t-51-72t-80-41t-107-13q-76 0-143 25t-128 72V462q71-42 150-60t160-18q74 0 142 20t121 61t83 102t31 142q0 139-70 223t-202 122v4z"/>`),
		g.Group(children),
	)
}