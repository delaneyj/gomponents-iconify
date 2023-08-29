package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reliability(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 384h-384q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h384q26 0 45 19t19 45.5t-19 45t-45 18.5zm-576 640h-320q-27 0-45.5-18.5T.193 960V640q0-26 18.5-45t45.5-19h320q35 0 54 30l-230 231l-54-54q-15-16-37-16t-37.5 15.5t-15.5 37.5t15 38l86 86q17 17 43 14q26 3 43-14l197-197v213q0 27-19 45.5t-45 18.5zm0-576h-320q-27 0-45.5-18.5T.193 384V64q0-26 18.5-45t45.5-19h320q35 0 54 30l-230 231l-54-54q-15-16-37-16t-37.5 15.5t-15.5 37.5t15 38l86 86q17 17 43 14q26 3 43-14l197-197v213q0 27-19 45.5t-45 18.5zm192 320h384q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19z"/>`),
		g.Group(children),
	)
}