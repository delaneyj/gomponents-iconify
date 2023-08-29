package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h341zM123 282q-38-37-38-90t38-91L93 71q-50 50-50 121t50 121zm90.5-5q35.5 0 60.5-25t25-60t-25-60t-60.5-25t-60.5 25t-25 60t25 60t60.5 25zM334 313q50-50 50-121T334 71l-30 31q37 37 37 90t-37 91zM213.5 149q17.5 0 30 12.5T256 192t-12.5 30.5t-30 12.5t-30-12.5T171 192t12.5-30.5t30-12.5z"/>`),
		g.Group(children),
	)
}