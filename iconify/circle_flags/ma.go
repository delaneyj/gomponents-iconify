package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMa0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#496e2d" d="M407.3 210H291.7L256 100.3L220.3 210H104.7l93.5 68l-35.7 109.8L256 320l93.5 68l-35.7-110zm-183 59.5l12.2-37.1h39l12.1 37.1l-31.6 23l-31.6-23zm44-59.4h-24.6l12.3-37.9zm38.3 45.7l-7.7-23.4h39.9zM213 232.4l-7.7 23.4l-32.2-23.4zm-8.3 97.3l12.3-38l20 14.5zm70.1-23.4l20-14.5l12.3 37.9z"/></g>`),
		g.Group(children),
	)
}