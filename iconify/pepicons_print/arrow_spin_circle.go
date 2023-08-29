package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSpinCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g transform="translate(3 3)"><circle cx="12" cy="11.5" r="6.5" opacity=".2"/><path d="M13.374 5.038a.5.5 0 0 1-.563.827A5 5 0 1 0 15 10a.5.5 0 0 1 1 0a6 6 0 1 1-2.626-4.962Z"/><path d="M12.769 11.585a.5.5 0 1 1-.539-.842l3.482-2.227a.5.5 0 1 1 .539.842l-3.482 2.227Z"/><path d="M17.947 12.114a.5.5 0 0 1-.913.407l-1.509-3.38a.5.5 0 1 1 .914-.408l1.508 3.381Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}