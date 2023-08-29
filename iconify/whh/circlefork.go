package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlefork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M929 808L719 631q23-49 13.5-103.5T684 433L581 332q-31-31-64-50t-51-22l-18-4q0 4 2 11.5t18.5 32.5t43.5 52l106 132q8 5 6.5 25T608 544L480 416q-23-23-55-39t-52-20l-21-5q1 8 4 22t19.5 48t40.5 58l128 128q-16 16-34.5 18t-25.5-8L352 512q-29-29-53-45t-33-17l-10-2q1 7 4 20t22 48t50 65l102 102q39 40 93.5 49.5T630 719l179 210q-134 95-297 95q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 163-95 296z"/>`),
		g.Group(children),
	)
}