package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShAc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsShAc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsShAc0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#6da544" d="m320 144l48-80l48 80z"/><circle cx="368" cy="144" r="48" fill="#acabb1"/><path fill="#338af3" d="M320 144v77c0 36 48 48 48 48s48-12 48-48v-77z"/><rect width="32" height="128" x="288" y="128" fill="#ff9811" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#ff9811" rx="16" ry="16"/><path fill="#6da544" d="m368 160l-48 67c2 11 9 19 16 26l32-45l32 45c8-7 14-15 16-26z"/></g>`),
		g.Group(children),
	)
}