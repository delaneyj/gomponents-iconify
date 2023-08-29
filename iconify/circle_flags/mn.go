package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMn0)"><path fill="#a2001d" d="M0 0h167l84.9 45L345 0h167v512H345l-87.7-48.1L167 512H0z"/><path fill="#0052b4" d="M167 0h178v512H167z"/><g fill="#ffda44"><path d="M122.4 256h22.3v89h-22.3zm-89 0h22.3v89H33.4z"/><circle cx="89" cy="289.4" r="22.3"/><circle cx="89" cy="211.5" r="11.1"/><path d="M66.8 322.8h44.5V345H66.8zm0-89h44.5V256H66.8zM89 133.5l8 24.2h25.4l-20.6 15l7.9 24.3L89 182l-20.6 15l7.9-24.3l-20.6-15h25.5z"/></g></g>`),
		g.Group(children),
	)
}