package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsWf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsWf0)"><path fill="#d80027" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0h256v256H0Z"/><path fill="#0052b4" d="M0 0h75v224H0Z"/><path fill="#d80027" d="M149 0h75v224h-75z"/><path fill="#eee" d="m384 232l-72-72h144zm-24 24l-72-72v144zm24 24l-72 72h144zm24-24l72-72v144z"/></g>`),
		g.Group(children),
	)
}