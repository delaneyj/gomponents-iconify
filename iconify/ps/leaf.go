package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 512q21 0 21-21v-66q63-7 103-45q47-44 47-124q0-23-2-34q4-13-3-21l-2-3q-10-42-32-79t-49-61t-42.5-36T189 3l-2-1l-5-2h-12l-5 2q-3 1-15 8.5T117.5 35T77 74t-38.5 55T12 198l-2 3q-7 8-3 21q-2 11-2 34q0 80 47 124q40 38 103 45v66q0 21 21 21zm21-452q36 28 60 59l-60 50V60zm0 164l84-66q6 11 17 49l-101 72v-55zm0 107l107-77v2q0 63-34 94q-29 25-73 34v-53zM155 60v109l-60-47q15-23 60-62zm-84 98l84 66v55L54 207q7-25 17-49zm11 192q-34-34-34-94v-2l107 77v51q-41-4-73-32z"/>`),
		g.Group(children),
	)
}