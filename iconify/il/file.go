package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M347 240q-9 0-16-7t-7-16V8l232 232H347zm-69-23q0 29 20 49t49 20h209v440q0 10-6 17t-17 7H23q-10 0-16-7t-7-17V31q0-9 7-16t16-7h255v209z"/>`),
		g.Group(children),
	)
}