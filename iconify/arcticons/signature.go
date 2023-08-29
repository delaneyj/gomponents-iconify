package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="23.762" cy="21.062" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.284" ry="6.182"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.55 27.237h1.03m-5.331 0H4.5a13.095 13.095 0 0 1 4.79-9.115c-.285 3.627-.284 6.08.754 9.115m.926-10.069a28.182 28.182 0 0 1 9.51-3.887"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.864 27.237h-1.2m.617 0a11.99 11.99 0 0 1 .557-12.577m13.442-1.24l3.876 1.46m-2.002-.73c3.322 5.492 2.065 11.52-.467 13.087h4.587m1.833-10.207l2.219 1.733m-1.228-.96c.752 3.488-.233 8.805-1.344 9.434h6.748v-.814M5.81 32.589l36.22-.03M9.062 34.868l29.848-.149M5.262 23.74l3.917-.716"/>`),
		g.Group(children),
	)
}