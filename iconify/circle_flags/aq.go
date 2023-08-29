package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAq0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="m135 343l-41-70l17-38l-40-51l-9-37l74 51l45-11l19-67l50-29l75 11l87 45l4 74l28 10v76l-53 94l-64 20l-59-14l15-25l-7-26l-8 7z"/></g>`),
		g.Group(children),
	)
}