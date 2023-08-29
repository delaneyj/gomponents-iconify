package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnitedNations(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUnitedNations0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUnitedNations0)"><path fill="#338af3" d="M0 0h512v512H0z"/><circle cx="256" cy="256" r="145" fill="#eee"/><circle cx="256" cy="256" r="111" fill="#338af3"/><path fill="#338af3" d="M76 76h360L256 256z"/><circle cx="256" cy="256" r="89" fill="#eee"/><circle cx="256" cy="256" r="69" fill="#338af3"/><path fill="#eee" d="M246 178h20v156h-20z"/><path fill="#eee" d="M334 246v20H178v-20z"/><path fill="#eee" d="m304 193.7l14.2 14.2l-110.3 110.3l-14.2-14.1z"/><path fill="#eee" d="m318.2 304l-14.1 14.2l-110.4-110.3l14.2-14.2z"/><circle cx="256" cy="256" r="44" fill="#eee"/><circle cx="256" cy="256" r="22" fill="#338af3"/><ellipse cx="256" cy="412" fill="#eee" rx="44" ry="40"/><path fill="#338af3" d="m256 407l-78 63h156z"/></g>`),
		g.Group(children),
	)
}