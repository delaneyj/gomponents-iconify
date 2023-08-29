package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavIconGridA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5.217a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zm8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zm-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zM12 14.608a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zm8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zm-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zM12 23.999a2.608 2.608 0 1 1 0-5.216A2.608 2.608 0 0 1 12 24zm8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216zm-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216z"/>`),
		g.Group(children),
	)
}