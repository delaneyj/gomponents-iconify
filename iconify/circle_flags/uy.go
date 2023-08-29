package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUy0)"><path fill="#338af3" d="M0 256L256 0h256v55.7l-20.7 34.5l20.7 32.2v66.8l-21.2 32.7L512 256v66.8l-24 31.7l24 35.1v66.7l-259.1 28.3L0 456.3v-66.7l27.1-33.3L0 322.8z"/><path fill="#eee" d="M256 256h256v-66.8H236.9zm-19.1-133.6H512V55.7H236.9zM512 512v-55.7H0V512zM0 389.6h512v-66.8H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#ffda44" d="m222.6 149.8l-31.3 14.7l16.7 30.3l-34-6.5l-4.3 34.3l-23.6-25.2l-23.7 25.2l-4.3-34.3l-33.9 6.5l16.6-30.3l-31.2-14.7l31.2-14.7l-16.6-30.3l34 6.5l4.2-34.3l23.7 25.3L169.7 77l4.3 34.3l34-6.5l-16.7 30.3z"/></g>`),
		g.Group(children),
	)
}