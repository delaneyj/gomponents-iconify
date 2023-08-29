package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wizardhat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-26 0-45-19T0 959.5t19-45T64 896h192l142-615l-142 39h-32q-13 0-22.5-9.5T192 288t9.5-22.5T224 256L512 0q26 0 45 19t19 45l192 832h192q27 0 45.5 18.5T1024 960t-18.5 45.5T960 1024z"/>`),
		g.Group(children),
	)
}