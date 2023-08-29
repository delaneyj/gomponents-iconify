package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm222-199l-79-79l-173 148q19 2 30 2q122 0 222-71zM576 704h37l-25-25zM128 512q0 150 101.5 260T479 894L320 704h128l32-128H256l67-162l-124-124q-71 100-71 222zm384-384q-122 0-222 71l87 87l39-94h160l-80 213l43 43h165l-55 110l176 176q71-100 71-222q0-104-51.5-192.5t-140-140T512 128z"/>`),
		g.Group(children),
	)
}