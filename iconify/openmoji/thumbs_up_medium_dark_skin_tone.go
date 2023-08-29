package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpMediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" stroke="#a57939" d="m37.088 11.232l-2.922.645L19 29.21l-7.333 13.059l1.999 11.274l12.938 7.07l9.625-.041l14.042.208l2.395-1.57l.75-3.833l-.75-1.584l3.584-.083l1.333-1.667l.833-2.333l-1.583-3l1.583-.667l1.917-2.666l-.917-2.75l-2.5-1.167l1.167-2.167l.167-3.416l-2.584-1.5l-13.166-.5l-.834-.667l-.854-.597h-9.125l.063-1.736l3.916-9.584l1.667-4.583z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M35.236 59.6h-9c-8.321 0-14-6.68-14-15c.02-4.799 1.8-9.424 5-13m17.389 0h-4.389l7-18c.811-2.083-2.79-3.8-5-1l-15 19m20.212 3.5a3.786 3.786 0 0 0 4 3.5h13a3.531 3.531 0 1 0 0-7h-13a3.786 3.786 0 0 0-4 3.5zm.857 21.337a3.786 3.786 0 0 0 4 3.5h7a3.531 3.531 0 1 0 0-7h-7a3.786 3.786 0 0 0-4 3.5zm-5.069-14a3.786 3.786 0 0 0 4 3.5h19a3.531 3.531 0 1 0 0-7h-19a3.786 3.786 0 0 0-4 3.5z"/><path stroke-linecap="round" stroke-linejoin="round" d="M35.236 49.437a3.786 3.786 0 0 0 4 3.5h15a3.531 3.531 0 1 0 0-7h-15a3.786 3.786 0 0 0-4 3.5z"/><path stroke-miterlimit="10" d="M14.136 36.263a19.457 19.457 0 0 0-1.906 7.839c0 8.56 4.625 15.5 15.125 15.5"/></g>`),
		g.Group(children),
	)
}