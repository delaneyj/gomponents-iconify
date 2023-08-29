package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mxentry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.356 384h-896q-26 0-45-19t-19-45V64q0-27 19-45.5t45-18.5h896q26 0 45 18.5t19 45.5v256q0 26-19 45t-45 19zm-767.5-256q-26.5 0-45.5 18.5t-19 45t19 45.5t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5zm639.5 0h-384q-26 0-45 18.5t-19 45t19 45.5t45 19h384q27 0 45.5-19t18.5-45.5t-18.5-45t-45.5-18.5zm-383 320l18 18l-72 71q-5 6-8 15.5t-3 16.5v7q-1 26 11 39l72 71l-275 275V448h257zm348 569q-14 7-29 7h-512q-15 0-29-7l285-286zm-222-569h257v513l-385-385z"/>`),
		g.Group(children),
	)
}