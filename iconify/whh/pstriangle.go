package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pstriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm32-815q-19-18-32-18t-32 18L192 706q0 22 4.5 35.5t15 18.5t19 6.5T256 768h512q16 0 25.5-1.5t19.5-7t14.5-19T832 706zm-32 79l256 416H256z"/>`),
		g.Group(children),
	)
}