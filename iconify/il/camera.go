package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 264q29 0 49 20t20 50t-20 49t-49 20t-49-20t-21-49t21-50t49-20zm347-162q10 0 17 6t6 17v440q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V125q0-10 7-17t16-6h139l16-56q3-16 16-27t29-10h295q17 0 30 10t15 27l16 56h139zM371 496q33 0 63-13t51-35t35-52t13-62q0-34-13-64t-35-51t-51-35t-63-13.5t-63 13.5t-52 35t-35 51t-13 64q0 33 13 62t35 52t52 35t63 13z"/>`),
		g.Group(children),
	)
}