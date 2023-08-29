package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCk0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="m345 256l4.8 14.6H365l-12.4 9l4.7 14.6l-12.4-9l-12.4 9l4.8-14.6l-12.4-9h15.3zm-63 26l13.7 7l10.9-10.8l-2.4 15.1l13.6 7l-15.1 2.4l-2.4 15.1l-7-13.6l-15.1 2.4l10.8-10.9zm-26 63l14.6-4.7V325l9 12.4l14.6-4.8l-9 12.4l9 12.4l-14.6-4.7l-9 12.4v-15.3zm26 63l7-13.6l-10.8-10.9l15.1 2.4l7-13.6l2.4 15l15.1 2.5l-13.6 7l2.4 15l-10.9-10.8zm63 26l-4.7-14.5H325l12.4-9l-4.8-14.6l12.4 9l12.4-9l-4.7 14.6l12.4 9h-15.3zm63-26l-13.6-7l-10.9 10.9l2.4-15.2l-13.6-7l15-2.3l2.5-15.1l7 13.6l15-2.4l-10.8 10.9zm26-63l-14.5 4.8V365l-9-12.4l-14.6 4.7l9-12.4l-9-12.4l14.6 4.8l9-12.4v15.3zm-26-63l-7 13.7l10.9 10.9l-15.2-2.4l-7 13.6l-2.3-15.1l-15.1-2.4l13.6-7l-2.4-15.1l10.9 10.8z"/></g>`),
		g.Group(children),
	)
}