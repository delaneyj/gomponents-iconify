package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func To(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTo0)"><path fill="#d80027" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#d80027" d="M167 133.6v-33.4h-33.4v33.4h-33.4V167h33.4v33.3H167V167h33.3v-33.4z"/></g>`),
		g.Group(children),
	)
}