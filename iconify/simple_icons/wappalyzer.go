package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wappalyzer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 11.014v-.604L12 1.805L0 10.41v.603l12 8.606l12-8.605zM8.634 10.82l2.75 1.07l.016-.01l-1.526-1.967l.984-.72l2.695 1.116l.016-.011l-1.463-2.018l1.247-.913l2.6 3.85l-1.046.766l-2.797-1.157l-.012.008l1.593 2.038l-1.048.767l-5.26-1.903l1.251-.916zm14.418 1.488l.947.679v.603l-12 8.605L0 13.59v-.603l.947-.678l10.761 7.717l.292.21l.291-.21l10.762-7.717z"/>`),
		g.Group(children),
	)
}