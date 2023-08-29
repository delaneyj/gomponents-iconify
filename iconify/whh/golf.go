package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Golf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-30T68.5 912.5T0 800t69.5-112.5T256 606v194q0 40 28 68t68 28t68-28t28-68V578q37-2 64-2q139 0 257 30t186.5 81.5T1024 800t-68.5 112.5T769 994t-257 30zm128-256q-27 0-45.5 19T576 832.5t18.5 45t45 18.5t45.5-18.5t19-45.5t-18.5-45.5T640 768zm99-544L477 384q-12 0-20.5-9.5T448 352V32q0-13 8.5-22.5T477 0l262 160q1 1 5 3t5.5 3t5 3t4.5 3.5t3.5 4t3.5 4.5t1.5 5t.5 6q0 12-7 18t-22 14zM352 832q-13 0-22.5-9.5T320 800V32q0-13 9.5-22.5T352 0t22.5 9.5T384 32v768q0 13-9.5 22.5T352 832z"/>`),
		g.Group(children),
	)
}