package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaRadar0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaRadar1" d="M 224.97499,72.065141 A 111.96707,111.96707 0 0 1 207.18133,207.22135 111.96707,111.96707 0 0 1 72.025105,225.01501 111.96707,111.96707 0 0 1 19.856764,99.069471 111.96707,111.96707 0 0 1 128.00866,16.081608"/><path id="galaRadar2" stroke-opacity="1" d="M 128.00866,16.081615 V 112.05339"/><path id="galaRadar3" d="m 144.00394,128.04868 a 15.995295,15.995295 0 0 1 -15.99528,15.9953 15.995295,15.995295 0 0 1 -15.99529,-15.9953 15.995295,15.995295 0 0 1 15.99529,-15.99529 15.995295,15.995295 0 0 1 15.99528,15.99529 z"/><path id="galaRadar4" d="M 183.41795,96.058104 A 63.981179,63.981179 0 0 1 173.25016,173.29022 63.981179,63.981179 0 0 1 96.018049,183.45803 63.981179,63.981179 0 0 1 66.207591,111.48915 63.981179,63.981179 0 0 1 128.00866,64.067515"/></g>`),
		g.Group(children),
	)
}