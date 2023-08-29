package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointerRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M152 341h43q17 0 29.5-12.5T237 299h-85q-47 0-79-35t-28-83q5-41 37-68.5T156 85h81q18 0 30.5-12.5T280 43H156Q98 43 54 80T3 173q-8 67 37 117.5T152 341zm128-106h-64q-21 0-21 21t21 21h64q21 0 21-21t-21-21zm0-64h-64q-21 0-21 21t21 21h64q21 0 21-21t-21-21zm149-43q0-21-21-21H216v42h192q21 0 21-21z"/>`),
		g.Group(children),
	)
}