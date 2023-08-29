package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Au(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAu0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m154 300l14 30l32-8l-14 30l25 20l-32 7l1 33l-26-21l-26 21l1-33l-33-7l26-20l-14-30l32 8zm222-27h47l-38 27l15-44l14 44zm7-162l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zm57 67l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm-122 22l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm65 156l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zM0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/></g>`),
		g.Group(children),
	)
}