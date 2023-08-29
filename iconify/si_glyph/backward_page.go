package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackwardPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m.252 4.516l3.905-3.25c.276-.288.593-.335.871-.048c0 0-.037 2.19-.006 2.19c.679 0 1.083.113 1.7.314A6.329 6.329 0 0 1 8.93 4.99c1.276 1.123 2.093 2.402 2.093 4.213c0 1.182-.523.752-.742.339c-1.035-1.945-2.923-2.928-5.282-2.928c-.03 0 .003 2.146.003 2.146a.607.607 0 0 1-.87.047l-3.88-3.25a.752.752 0 0 1 0-1.041z"/><path d="M1.954 9.079v4.983h13.118V4.937H9.871c-.275-.312-.619-.717-1.219-.919H16v10.964H1.08V8.258l.874.821z"/></g>`),
		g.Group(children),
	)
}