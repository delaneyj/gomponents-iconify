package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorK0" d="m37.316 31.46l7.746-8.79a1 1 0 0 0-1.5-1.321L29.235 37.605V22.012a1 1 0 1 0-2 0v28a1 1 0 1 0 2 0V40.63l6.663-7.56l8.98 17.413a1 1 0 1 0 1.777-.917L37.316 31.46Z"/></defs><use href="#openmojiRegionalIndicatorK0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorK0"/></g>`),
		g.Group(children),
	)
}