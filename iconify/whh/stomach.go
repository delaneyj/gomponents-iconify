package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stomach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 960q-65 0-125.5-30T227 855.5T192 768q0-30-16-42.5t-32 0t-16 42.5v192q0 21-21.5 42.5T64 1024t-42.5-21.5T0 960V768q0-38 2.5-64t10-52T34 610.5t37.5-25T128 576q34 0 62 10t43.5 22t38.5 22t48 10q27 0 42-.5t42.5-3t45-8.5t38.5-16t34.5-27.5T544 544q15-45 22.5-99.5T575 355t1-99q-48 0-88-40t-40-88V64q0-21 21.5-42.5T512 0t42.5 21.5T576 64v32q0 32 128 32q17 0 33.5-10T767 96t39-22t58-10q75 0 117.5 94.5T1024 410q0 131-45 236.5T855.5 820T673 924t-225 36z"/>`),
		g.Group(children),
	)
}