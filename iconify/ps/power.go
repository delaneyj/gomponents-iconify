package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M312 107q-19-9-28 10q-8 21 11 28q42 23 67 64t25 90q0 70-50 120t-121 50t-121-50t-50-120q0-47 25-88t67-64q18-10 11-30q-13-14-30-8q-53 27-84 78T3 299q0 88 62.5 150.5T216 512t150.5-62.5T429 299q0-61-31.5-113T312 107zm-75 149V21q0-21-21-21t-21 21v235q0 21 21 21t21-21z"/>`),
		g.Group(children),
	)
}