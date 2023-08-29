package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsAk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUsAk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsAk0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m336.8 199.7l104-75.1h-128l104 75.1l-40-122.1zM57 231.7l52-37.4H45l52 37.4l-20-60.9zm85.5 31.7l52-37.4h-64l52 37.4l-20-60.9zM185 309l52-37.4h-64l52 37.4l-20-60.9zm43 47.6l52-37.4h-64l52 37.4l-20-60.9zm-5.6 67.1l52-37.4h-64l52 37.4l-20-60.9zM356 402.2l52-37.4h-64l52 37.4l-20-60.9zm-37.4 53.2l52-37.4h-64l52 37.4l-20-60.9z"/></g>`),
		g.Group(children),
	)
}