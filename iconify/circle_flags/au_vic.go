package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuVic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAuVic0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuVic0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="m313 267l7 15l15-4l-7 15l13 10l-16 3v16l-13-10l-12 10v-16l-16-3l13-10l-7-15l16 4zm66-61l7 15l15-4l-7 15l13 10l-16 4v16l-13-10l-12 10v-16l-16-4l13-10l-7-15l16 4zm1 146l40-29h-50l40 29l-15-48Zm-1 29l5 15l14-7l-7 15l15 5l-15 5l7 15l-14-7l-5 15l-6-15l-14 7l7-15l-15-5l15-5l-7-15l14 7zm57-114l-8 14h-16l8 14l-8 14h16l8 13l8-13h16l-8-14l8-14h-16z"/><path fill="#ffda44" d="M370 93v8h-8v17h8v15a25 25 0 0 0-34 37v15h85v-15a25 25 0 0 0-34-37v-15h8v-17h-8v-8z"/></g>`),
		g.Group(children),
	)
}