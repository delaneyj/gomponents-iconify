package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photosphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M875 710q-62 86-157.5 136T512 896t-205.5-50T149 710Q0 636 0 448t149-262q62-86 157.5-136T512 0t205.5 50T875 186q149 74 149 262T875 710zM512 192q-88 0-154 7t-123.5 24.5t-93.5 46T84.5 343T64 448q0 88 41.5 142T229 671l141-141q18-19 46.5-19t46.5 19l97 97l162-161q18-19 46.5-19t46.5 19l112 112q33-52 33-130q0-60-20.5-105T883 269.5t-93.5-46T666 199t-154-7zM256 512q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}