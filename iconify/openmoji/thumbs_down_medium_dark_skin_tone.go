package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDownMediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" stroke="#a57939" d="m37.333 57.207l-1.667-4.583l-3.916-9.584l-.063-1.736h9.125l.854-.597l.834-.667l13.166-.5l2.584-1.5l-.167-3.416l-1.167-2.167l2.5-1.167l.917-2.75l-1.917-2.666l-1.583-.667l1.583-3l-.833-2.333l-1.333-1.667l-3.584-.083l.75-1.584l-.75-3.833l-2.395-1.57l-14.042.208l-9.625-.041l-12.938 7.07l-1.999 11.274L19 42.707L34.166 60.04l2.922.645z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M17.236 40.316c-3.2-3.576-4.98-8.2-5-13c0-8.319 5.679-15 14-15h9m-18 28l15 19c2.21 2.8 5.811 1.084 5-1l-7-18h4.389m6.823 0h13a3.53 3.53 0 1 0 0-7h-13a3.786 3.786 0 0 0-4 3.5a3.786 3.786 0 0 0 4 3.5zm.857-21.336h7a3.53 3.53 0 1 0 0-7h-7a3.786 3.786 0 0 0-4 3.5a3.786 3.786 0 0 0 4 3.5zm-5.069 14h19a3.53 3.53 0 1 0 0-7h-19a3.786 3.786 0 0 0-4 3.5a3.786 3.786 0 0 0 4 3.5z"/><path stroke-linecap="round" stroke-linejoin="round" d="M39.236 25.98h15a3.53 3.53 0 1 0 0-7h-15a3.786 3.786 0 0 0-4 3.5a3.786 3.786 0 0 0 4 3.5z"/><path stroke-miterlimit="10" d="M27.355 12.315c-10.5 0-15.125 6.94-15.125 15.5c.08 2.717.73 5.388 1.906 7.839"/></g>`),
		g.Group(children),
	)
}