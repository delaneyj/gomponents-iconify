package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Networksignalalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.84 1024q-26.5 0-45.5-18.5t-19-45.5V256q0-26 19-45t45.5-19t45 19t18.5 45v704q0 27-18.5 45.5t-45 18.5zm-256 0q-26.5 0-45.5-18.5t-19-45.5V512q0-26 19-45t45.5-19t45 19t18.5 45v448q0 27-18.5 45.5t-45 18.5zm-256 0q-26.5 0-45.5-18.5t-19-45.5V768q0-26 19-45t45.5-19t45 19t18.5 45v192q0 27-18.5 45.5t-45 18.5zm-192.5-784v752q0 13-9 22.5t-22.5 9.5t-23-9.5t-9.5-22.5V240l-182-183q-10-10-10-23.5t10-23.5t23.5-10t23.5 10l135 135V33q0-14 9.5-23t23-9t22.5 9t9 23v112l136-135q10-10 23.5-10t23.5 10t10 23.5t-10 23.5z"/>`),
		g.Group(children),
	)
}