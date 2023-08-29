package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capacitator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 384H704v192q0 26-19 45t-45.5 19t-45-19t-18.5-45V64q0-26 18.5-45t45-19T685 19t19 45v192h256q26 0 45 18.5t19 45t-19 45.5t-45 19zM383.5 640q-26.5 0-45-19T320 576V384H64q-27 0-45.5-19T0 319.5t18.5-45T64 256h256V64q0-26 18.5-45t45-19T429 19t19 45v512q0 26-19 45t-45.5 19z"/>`),
		g.Group(children),
	)
}