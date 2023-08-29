package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 61h-31L97 19h330q17 0 29.5 12.5T469 61v330l-42-43v-31h-31l-43-42h74v-86h-86v74l-42-43v-31h-31l-43-42h74V61h-86v74l-42-43V61zm170 0v86h86V61h-86zM27 3l458 458l-27 27l-43-43H85q-17 0-29.5-12.5T43 403V73L0 30zm186 241v31h31zM85 116v31h31zm86 287v-86H85v86h86zm0-128v-74l-12-12H85v86h86zm128 128v-74l-12-12h-74v86h86zm42 0h31l-31-31v31z"/>`),
		g.Group(children),
	)
}