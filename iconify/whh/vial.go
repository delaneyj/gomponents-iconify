package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M480 64h-32v768q0 80-56 136t-136 56t-136-56t-56-136V64H32q-13 0-22.5-9.5T0 32T9.5 9.5T32 0h448q13 0 22.5 9.5T512 32t-9.5 22.5T480 64zM288 832q13 0 22.5-9.5T320 800t-9.5-22.5T288 768t-22.5 9.5T256 800t9.5 22.5T288 832zm-96.5-128q26.5 0 45.5-19t19-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45t18.5 45.5t45 19zM384 64H128v320h256V64zm-32 448q-13 0-22.5 9.5T320 544t9.5 22.5T352 576t22.5-9.5T384 544t-9.5-22.5T352 512z"/>`),
		g.Group(children),
	)
}