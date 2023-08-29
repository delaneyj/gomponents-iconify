package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMs0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#338af3" d="M289.4 133.6V256l78 40.4l77.9-40.4V133.6h-156z"/><path fill="#a2001d" d="M289.4 256c0 59.6 77.9 78 77.9 78s78-18.4 78-78h-156z"/><path fill="#333" d="M400.7 189.2h-22.3V167h-22.2v22.2h-22.3v22.3h22.3v66.8h22.2v-66.8h22.3z"/></g>`),
		g.Group(children),
	)
}