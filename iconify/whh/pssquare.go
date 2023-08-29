package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pssquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm256-704q0-27-19-45.5T704 256H320q-26 0-45 18.5T256 320v384q0 26 19 45t45 19h384q27 0 45.5-19t18.5-45V320zm-448 0h384v384H320V320z"/>`),
		g.Group(children),
	)
}