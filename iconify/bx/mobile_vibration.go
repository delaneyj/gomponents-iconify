package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileVibration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.535 2.808c-.756-.756-2.072-.756-2.828 0l-9.899 9.899a2.001 2.001 0 0 0 0 2.828l5.657 5.657c.378.378.88.586 1.414.586s1.036-.208 1.414-.586l9.899-9.899c.378-.378.586-.88.586-1.414s-.208-1.036-.586-1.414l-5.657-5.657zm-5.656 16.97v1v-1l-5.657-5.657l9.899-9.899l5.657 5.657l-9.899 9.899z"/><circle cx="9" cy="15" r="1" fill="currentColor"/><path fill="currentColor" d="m15.707 21.707l-1.414-1.414l6-6l1.414 1.415zM8.293 2.293l1.414 1.414l-6 6l-1.414-1.415z"/>`),
		g.Group(children),
	)
}