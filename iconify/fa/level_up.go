package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1018 475q-18 37-58 37H768v864q0 14-9 23t-23 9H32q-21 0-29-18q-8-20 4-35l160-192q9-11 25-11h320V512H320q-40 0-58-37q-17-37 9-68L591 23q18-22 49-22t49 22l320 384q27 32 9 68z"/>`),
		g.Group(children),
	)
}