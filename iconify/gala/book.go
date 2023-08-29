package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaBook0"><g id="galaBook1" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="4.234" transform="translate(.005 .01) scale(3.77938)"><g id="galaBook2" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.234" transform="translate(4.233)"><path id="galaBook3" stroke-opacity="1" d="M 55.033333,63.5 12.7,63.499999"/><path id="galaBook4" d="M 4.2333336,-55.033333 A 8.4666662,8.4666662 0 0 1 12.7,-63.499999" transform="scale(1 -1)"/><path id="galaBook5" d="M 4.2333336,55.033333 A 8.4666662,8.4666662 0 0 1 12.7,46.566667"/><path id="galaBook6" stroke-opacity="1" d="m 55.033333,46.566667 -42.333332,-1e-6"/><path id="galaBook7" d="M 4.2333336,12.7 A 8.4666662,8.4666662 0 0 1 12.7,4.2333336"/><path id="galaBook8" stroke-opacity="1" d="m 55.033333,4.233334 -42.333332,-1e-6"/><path id="galaBook9" stroke-opacity="1" d="m 4.2333334,12.7 2e-7,42.333333"/><path id="galaBooka" stroke-opacity="1" d="M 12.7,46.566666 V 4.230783"/><path id="galaBookb" stroke-opacity="1" d="M 55.033333,4.2333348 V 46.566667"/><path id="galaBookc" stroke-opacity="1" d="m 55.033333,46.566667 c -2.662642,5.559118 -2.809222,11.198865 0,16.933333"/></g></g></g>`),
		g.Group(children),
	)
}