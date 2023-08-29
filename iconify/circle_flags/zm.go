package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsZm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsZm0)"><path fill="#496e2d" d="M0 0h512v256L256 512H0z"/><path fill="#ff9811" d="M473 167h-66.7a22.3 22.3 0 0 0-44.6 0H295a23 23 0 0 0 23 22.2h-.8c0 12.3 10 22.3 22.3 22.3c0 12.3 10 22.2 22.2 22.2h44.6c12.3 0 22.2-10 22.2-22.2c12.3 0 22.3-10 22.3-22.3h-.8a23 23 0 0 0 23-22.2z"/><path fill="#333" d="M341.3 256h85.4l21.1 126.3L426.7 512h-85.4l-23.5-128z"/><path fill="#d80027" d="M256 256h85.3v256H256z"/><path fill="#ff9811" d="M426.7 256H512v256h-85.3z"/></g>`),
		g.Group(children),
	)
}