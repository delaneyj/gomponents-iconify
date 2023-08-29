package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusBat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.211 20.568a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858Zm-.571-9.429h1.143m-.572 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571H20.64m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.852 3.576H16.64m.571 0v-2.571m-3.838 1.218l-.809-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.572M12.564 13.3l.809-.808m-.405.404l1.819 1.819M8.813.861v4.012L5.665 3.824A3.7 3.7 0 0 0 .789 7.339v2.719s2.485-1.043 4.076 2.619c2.038-2.038 4.075 0 4.075 2.038M14.163.861v4.012l3.147-1.049a3.705 3.705 0 0 1 4.877 3.515"/><path d="M8.813 3.536A4.667 4.667 0 0 1 11.488 2.2c1.014.14 1.954.61 2.675 1.338"/></g>`),
		g.Group(children),
	)
}