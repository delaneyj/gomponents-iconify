package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 4v11.469c-1.75 1.27-3 3.207-3 5.531c0 3.855 3.145 7 7 7s7-3.145 7-7c0-2.324-1.25-4.262-3-5.531V14h2v-2h-2v-2h2V8h-2V6h2V4zm2 2h4v10.406l.5.282c1.496.867 2.5 2.46 2.5 4.312c0 2.773-2.227 5-5 5s-5-2.227-5-5c0-1.852 1.004-3.445 2.5-4.313l.5-.28zm1 2v10.188c-1.16.414-2 1.511-2 2.812a3 3 0 0 0 6 0c0-1.3-.84-2.398-2-2.813V8z"/>`),
		g.Group(children),
	)
}