package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsNu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNu0)"><path fill="#ffda44" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32l42-16l41 16h45l-8-16l8-15v-14l-16-42l16-41V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#0052b4" d="M128 256v-83l83 83zm128-45l-83-83h83z"/><path fill="#d80027" d="m128 128l128 128v-31l-97-97Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Z"/><circle cx="64" cy="64" r="48" fill="#0052b4"/><path fill="#ffda44" d="m50 198l14-44l15 44l-38-27h47zM162 86l14-44l15 44l-38-27h47ZM64 17l28 86l-73-53h90l-73 53Z"/></g>`),
		g.Group(children),
	)
}