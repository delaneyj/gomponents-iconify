package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsPa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPa0)"><path fill="#eee" d="M0 0h256l256 256v256H256L0 256z"/><path fill="#0052b4" d="M0 256v256h256V256z"/><path fill="#d80027" d="M256 0h256v256H256z"/><path fill="#0052b4" d="m152.4 89l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51L82.2 140h53.6z"/><path fill="#d80027" d="m359.6 289.4l16.6 51h53.6L386.4 372l16.6 51l-43.4-31.5l-43.4 31.6l16.6-51l-43.4-31.6H343z"/></g>`),
		g.Group(children),
	)
}