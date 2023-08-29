package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deathstarbuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 576H768v64h128v23q-167 41-384 41q-137 0-256-17v-47H128v-64H64v64H46q-24-8-32-11Q0 569 0 512q0-104 40.5-199t109-163.5T313 40.5T512 0q95 0 182 33.5T850 128H704v64h128v64h64v64H768v64h128v64h124q4 33 4 64H896v64zM443 143.5Q389 115 323.5 137t-103 81.5T195 341t66 91.5t119.5 6.5t103-81.5T509 235t-66-91.5zM352 320q-13 0-22.5-9.5T320 288t9.5-22.5T352 256t22.5 9.5T384 288t-9.5 22.5T352 320zM128 708v60h384q219 0 384-60v60h-64v64H704v64h64v58l-6 3l-5 3h-53v26q-93 38-192 38q-159 0-288-88.5T38 704h79q2 1 5.5 2t5.5 2z"/>`),
		g.Group(children),
	)
}