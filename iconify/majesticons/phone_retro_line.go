package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneRetroLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 4.842C15.976 4.337 14.146 4 12 4c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 11l-3-2V4.842zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 11l3-2V4.842zM10 7v3m0 0l-5.414 5.414A2 2 0 0 0 4 16.828V18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-1.172a2 2 0 0 0-.586-1.414L14 10m-4 0h4m0 0V7"/><circle cx="12" cy="15" r="2"/></g>`),
		g.Group(children),
	)
}