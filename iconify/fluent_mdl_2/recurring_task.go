package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecurringTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1152v384h-384v-128h190q-45-60-112-94t-142-34q-59 0-111 20t-95 55t-70 85t-38 107l-127-22q14-81 54-149t98-118t133-78t156-28q91 0 174 35t146 102v-137h128zm-448 768q58 0 111-20t95-55t70-85t38-107l127 22q-14 81-54 149t-98 118t-133 78t-156 28q-91 0-174-35t-146-102v137h-128v-384h384v128h-190q45 60 112 94t142 34zM576 1061l627-626l90 90l-717 718l-333-334l90-90l243 242zm-448 347h912q-11 31-18 63t-10 65H0V0h1536v1012q-33 3-65 10t-63 18V128H128v1280z"/>`),
		g.Group(children),
	)
}