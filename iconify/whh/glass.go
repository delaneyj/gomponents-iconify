package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576.914 575v320h128q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5h-384q-26 0-45-18.5t-19-45t19-45.5t45-19h128V575l-439-479q-9-16-9-32q0-27 19-45.5t45-18.5h896q26 0 45 18.5t19 45.5q0 16-9 32zm-512-511l239 255h87l7-13l-73-131q-6-12-2.5-24.5t15-19t24-3t19.5 14.5l66 119q18-7 34-7q31 0 55.5 18t34.5 46h151l239-255h-896z"/>`),
		g.Group(children),
	)
}