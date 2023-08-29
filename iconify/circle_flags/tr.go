package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTr0)"><path fill="#d80027" d="M0 0h512v512H0z"/><g fill="#eee"><path d="m245.5 209.2l21 29l34-11.1l-21 29l21 28.9l-34-11.1l-21 29V267l-34-11.1l34-11z"/><path d="M188.2 328.3a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/></g></g>`),
		g.Group(children),
	)
}