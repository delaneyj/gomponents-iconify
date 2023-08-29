package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m444 82l-73 253s-83-58-144-52c-60 6-141 36-141 36l74-255s153-82 284 18zm-34 277l73-251c22 15 61 42 124 50c63 7 161-33 161-33l-73 253c-11 4-57 30-141 35c-85 5-144-54-144-54zM0 618l73-255s82-38 151-33c68 4 123 42 134 51l-73 255s-84-55-142-50S34 601 0 618zm324 40l73-253s78 51 129 49c51-1 86-4 155-31h1c-5 13-73 251-73 251s-80 38-154 34c-73-3-117-45-131-50z"/>`),
		g.Group(children),
	)
}