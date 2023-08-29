package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func King(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 832q26 0 45 19t19 45v64q0 27-19 45.5t-45 18.5H64q-26 0-45-18.5T0 960v-64q0-26 19-45t45-19q58 0 106-53t70-140q-101-6-170.5-79.5T0 384q0-106 75-181t181-75q32 0 64 8V64q0-26 19-45t45.5-19t45 19T448 64v72q32-8 64-8q106 0 181 75t75 181q0 102-69.5 175.5T528 639q22 87 70 140t106 53z"/>`),
		g.Group(children),
	)
}