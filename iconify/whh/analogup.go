package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024q-91 0-174-35.5T131 893T35.5 750T0 576q0-126 65-232t173-163l-31 50q-10 13-10.5 34t11 38t31.5 17l418-1q21 0 32-17t10.5-38t-10.5-33l-30-50q108 58 172.5 163.5T896 576q0 91-35.5 174T765 893t-143 95.5t-174 35.5zM297 256q-9-13-9-30.5t9-30.5L427 13q9-13 21.5-13T470 13l129 182q9 12 9 30t-9 31H297z"/>`),
		g.Group(children),
	)
}