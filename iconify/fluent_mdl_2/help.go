package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1920h128v128H896v-128zM960 0q79 0 152 20t138 58t117 91t90 117t58 137t21 153q0 84-22 152t-58 124t-82 105t-94 93t-94 89t-82 95t-58 108t-22 130v192H896v-192q0-84 22-152t58-124t82-104t94-93t94-90t82-95t58-108t22-130q0-93-35-174t-96-142t-142-96t-175-36q-93 0-174 35t-142 96t-96 142t-36 175H384q0-79 20-152t58-138t91-117t117-90t137-58T960 0z"/>`),
		g.Group(children),
	)
}