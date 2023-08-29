package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsPr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPr0)"><path fill="#eee" d="m27 63.3l485 39.1v102.4L477.3 259l34.7 48.2v102.4L27.4 446.9z"/><path fill="#d80027" d="m0 0l51.2 102.4H512V0zm0 512h512V409.6H51.2zm180-204.8h332V204.8H180z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#eee" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/></g>`),
		g.Group(children),
	)
}