package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fluxbb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M672 960h-64q-13 0-22.5-9.5T576 928v-64q0-13 9.5-22.5T608 832h32V192h-32q-13 0-22.5-9.5T576 160V96q0-13 9.5-22.5T608 64h64q40 0 68 28t28 68v704q0 40-28 68t-68 28zm-256 64h-64q-13 0-22.5-9.5T320 992V32q0-13 9.5-22.5T352 0h64q13 0 22.5 9.5T448 32v960q0 13-9.5 22.5T416 1024zm-256-64H96q-40 0-68-28T0 864V160q0-40 28-68t68-28h64q13 0 22.5 9.5T192 96v64q0 13-9.5 22.5T160 192h-32v640h32q13 0 22.5 9.5T192 864v64q0 13-9.5 22.5T160 960z"/>`),
		g.Group(children),
	)
}