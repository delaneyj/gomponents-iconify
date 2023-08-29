package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5ZM9.128 2A3.496 3.496 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2h6.628Zm8.53 7.586l-7.072 7.07l-4.242-4.242L7.758 11l2.828 2.829l5.657-5.657l1.414 1.414Z"/>`),
		g.Group(children),
	)
}