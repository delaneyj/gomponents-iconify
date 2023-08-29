package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm64-891v125q0 26-18.5 45T512 322t-45.5-19t-18.5-45V133q-120 20-207 107T134 448h122q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19H134q20 120 107 207t207 107V768q0-26 18.5-44.5T512 705t45.5 18.5T576 768v122q120-20 207-107t107-207H767q-26 0-44.5-19T704 511.5t18.5-45T767 448h123q-20-121-107-208T576 133z"/>`),
		g.Group(children),
	)
}