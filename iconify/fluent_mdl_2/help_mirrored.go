package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1920H896v128h128v-128zM960 0q-79 0-152 20T670 78t-117 91t-90 117t-58 137t-21 153q0 84 22 152t58 124t82 105t94 93t94 89t82 95t58 108t22 130v192h128v-192q0-84-22-152t-58-124t-82-104t-94-93t-94-90t-82-95t-58-108t-22-130q0-93 35-174t96-142t142-96t175-36q93 0 174 35t142 96t96 142t36 175h128q0-79-20-152t-58-138t-91-117t-117-90t-137-58T960 0z"/>`),
		g.Group(children),
	)
}