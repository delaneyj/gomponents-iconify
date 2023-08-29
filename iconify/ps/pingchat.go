package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pingchat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M143 330h139q0 56-40.5 95.5T143 465q-57 0-98-39.5T4 330q0-55 41-95t98-40v135zm322-193q0 55-40 93.5T329 269H193V130q2-53 41.5-89.5T329 4q56 0 96 39t40 94zm-116 66q0-8-6-14t-15-6q-8 0-14 6t-6 14t6 14t14 6q9 0 15-6t6-14zm0-134q0-8-6-14t-15-6q-8 0-14 6t-6 14v3l6 79q0 15 14 15q15 0 15-15l6-79v-3z"/>`),
		g.Group(children),
	)
}