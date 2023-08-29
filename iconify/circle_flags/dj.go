package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsDj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDj0)"><path fill="#338af3" d="M0 0h512v256l-153.2 35.7L210 256z"/><path fill="#6da544" d="M210 256h302v256H0z"/><path fill="#eee" d="M0 0v512l256-256z"/><path fill="#d80027" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/></g>`),
		g.Group(children),
	)
}