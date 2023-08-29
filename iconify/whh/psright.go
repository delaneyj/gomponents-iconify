package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M52 354L347 74q22-23 71-48.5T497 0h377q62 0 106 43.5t44 105.5v598q0 61-44 105t-106 44H497q-30 0-82-25t-68-50L52 542Q0 502 0 448t52-94z"/>`),
		g.Group(children),
	)
}