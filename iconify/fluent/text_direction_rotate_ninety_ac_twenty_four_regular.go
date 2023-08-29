package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateNinetyAcTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M18.78 17.72a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72v-3.186a.75.75 0 0 1 1.5 0v3.185l.72-.72a.75.75 0 0 1 1.06 0zm-9.5 1.06a.75.75 0 1 0-1.06-1.06l-.72.72V3.75a.75.75 0 0 0-1.5 0v14.69l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2zM20 10.75a.75.75 0 0 1-.75.75h-3v1.25a.75.75 0 0 1-1.5 0V11.5h-4.5a.75.75 0 0 1 0-1.5h9a.75.75 0 0 1 .75.75zM18.5 3a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-.75.75h-1.25c-1.078 0-2.426-.212-3.53-.918c-1.15-.737-1.97-1.973-1.97-3.832a.75.75 0 0 1 1.5 0c0 1.341.555 2.105 1.28 2.568c.771.494 1.798.682 2.72.682h.5V3.75A.75.75 0 0 1 18.5 3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}