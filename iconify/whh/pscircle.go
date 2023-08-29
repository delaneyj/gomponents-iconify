package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pscircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-832q-87 0-160.5 42.5T235 351t-43 161t43 160.5T351.5 789T512 832t160.5-43T789 672.5T832 512t-43-161t-116.5-116.5T512 192zm0 576q-106 0-181-75t-75-181t75-181t181-75t181 75t75 181t-75 181t-181 75z"/>`),
		g.Group(children),
	)
}