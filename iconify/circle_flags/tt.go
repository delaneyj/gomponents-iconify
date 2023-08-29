package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTt0)"><path fill="#d80027" d="M0 110.2L110.2 0H512v401.8L401.8 512H0z"/><path fill="#eee" d="M110.2 0H63L0 63v47.2L401.8 512H449l63-63v-47.2z"/><path fill="#333" d="M512 512v-63L63 0H0v63l449 449z"/></g>`),
		g.Group(children),
	)
}