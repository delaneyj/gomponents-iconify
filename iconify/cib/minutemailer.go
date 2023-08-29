package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minutemailer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m7.416 3.604l4.605 4.98l-3.251 6.395l6.855-1.229l3.12 7.532L32 3.902zm-.843 10.781l1.276-1.047l-1.647.521l-.172-.24l.683-.661l-.891.359c-3.407 1.323-5.823 4.62-5.823 8.485a9.043 9.043 0 0 0 2.844 6.593A9.006 9.006 0 0 1 1.66 23.92c0-3.817 2.417-7.219 5.755-8.557l.423-1.02l-1 .437l-.281-.38zm5.818-2.625L14.522 8l12.531-2.932z"/>`),
		g.Group(children),
	)
}