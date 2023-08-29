package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4.5h7.5V12h-2V7.914L7.914 17.5H12v2H4.5V12h2v4.086L16.086 6.5H12v-2Z"/>`),
		g.Group(children),
	)
}