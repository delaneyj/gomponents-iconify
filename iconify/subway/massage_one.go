package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M78.8 56.1v78.8H0V450h433.2v-78.8H512V56.1H78.8zm354.4 39.4L295.4 233.3L157.5 95.5h275.7zm-39.4 315H39.4V193.9l39.4 39.4v137.8h315.1v39.4zm78.8-78.7H118.2V115.2l177.2 177.2l177.2-177.2v216.6z"/>`),
		g.Group(children),
	)
}