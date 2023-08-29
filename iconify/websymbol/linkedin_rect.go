package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinRect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 89v824q0 36-26 62t-62 26H88q-36 0-62-26T0 913V89q0-36 26-62T88 1h824q36 0 62 26t26 62zM668 801h129V554q0-89-41-135t-107-46q-47 0-79.5 21T521 441h-2l-6-59H402q3 114 3 134v285h128V560q0-21 4-33q20-49 66-49q65 0 65 91v232zM313.5 219.5Q295 201 265 201t-49 18.5t-19 46.5t18.5 46.5T263 331q32 0 51-18q18-18 18-44q0-31-18.5-49.5zM200 801h128V382H200v419z"/>`),
		g.Group(children),
	)
}