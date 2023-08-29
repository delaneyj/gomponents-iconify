package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1088 4q96 0 185 25t167 71t142 110t110 142t71 167t25 185q0 137-52 264t-150 225l-837 836l-90-90l836-837q79-79 122-182t43-216q0-117-45-221t-124-182t-182-123t-221-46q-108 0-190 32t-153 86t-134 122t-136 140h421v128H256V0h128v421q55-56 105-108t101-99t103-85t112-66t129-43t154-16z"/>`),
		g.Group(children),
	)
}