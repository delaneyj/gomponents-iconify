package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 576h-512q-26 0-45-18.5t-19-45t19-45.5t45-19h512q26 0 45 19t19 45.5t-19 45t-45 18.5zm0-384h-512q-26 0-45-18.5t-19-45t19-45.5t45-19h512q26 0 45 19t19 45.5t-19 45t-45 18.5zm-768 832h-128q-27 0-45.5-19t-18.5-45V832q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45v128q0 27-19 45.5t-45 18.5zm0-384h-128q-27 0-45.5-18.5T.193 576V448q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45v128q0 27-19 45.5t-45 18.5zm0-384h-128q-27 0-45.5-18.5T.193 192V64q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45v128q0 27-19 45.5t-45 18.5zm256 576h512q26 0 45 19t19 45.5t-19 45t-45 18.5h-512q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19z"/>`),
		g.Group(children),
	)
}