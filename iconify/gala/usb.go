package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaUsb0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaUsb1" stroke-opacity="1" d="M 48.042103,208.01896 V 112.07317"/><path id="galaUsb2" stroke-opacity="1" d="M 207.96306,208.01897 V 112.06638"/><path id="galaUsb3" stroke-opacity="1" d="M 80.026346,240.00316 H 175.97888"/><path id="galaUsb4" stroke-opacity="1" d="m 207.96306,208.01897 a 31.98419,31.98419 0 0 1 -31.98418,31.98419"/><path id="galaUsb5" stroke-opacity="1" d="m 48.042103,208.01897 a 31.98419,31.98419 0 0 0 31.984243,31.98419"/><path id="galaUsb6" d="M 48.035077,112.07316 A 15.992095,15.992095 0 0 1 64.027161,96.081066"/><path id="galaUsb7" d="M 207.96306,112.06637 A 15.992095,15.992095 0 0 0 191.97096,96.074276"/><path id="galaUsb8" stroke-opacity="1" d="M 64.034224,96.074284 H 191.97096"/><path id="galaUsb9" d="M 80.026346,32.112684 C 80.026347,23.280513 87.167829,16.000027 96,16"/><path id="galaUsba" d="M 175.97888,32.105897 C 175.97888,23.273697 168.8322,15.999986 160,16"/><path id="galaUsbb" stroke-opacity="1" d="M 96.018393,16 H 160"/><path id="galaUsbc" stroke-opacity="1" d="M 80.026346,32.105904 V 96.074281"/><path id="galaUsbd" stroke-opacity="1" d="M 175.97888,32.105904 V 96.074281"/><path id="galaUsbe" stroke-opacity="1" d="m 112.01048,16.113809 0,31.984193"/><path id="galaUsbf" stroke-opacity="1" d="M 143.99468,16.113809 V 48.098002"/></g>`),
		g.Group(children),
	)
}