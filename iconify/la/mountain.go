package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.012 3.021l-.912 1.66l-6.522 11.856l-1.916-1.916l-.66 1.098l-5.86 9.767L.235 27h31.284l-.598-1.395l-3-7l-.582-1.357l-2.068 2.068l-7.403-14.605l-.855-1.69zm-.073 4.282l3.04 5.996l-.774.664l-2.28-1.953l-2.279 1.953l-.93-.799l3.223-5.861zm-.013 7.34l2.28 1.953l1.702-1.46l3.2 6.315l.622 1.233l1.932-1.932L28.482 25H3.766c1.43-2.385 2.862-4.77 4.293-7.154l1.988 1.988l.642-1.166c.681-1.238 1.363-2.475 2.043-3.713l1.914 1.64l2.28-1.952z"/>`),
		g.Group(children),
	)
}