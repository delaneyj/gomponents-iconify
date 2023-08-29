package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="M40.184 12L26.13 26.859l3.729 3.413l7.51-7.944V38.82c0 10.524-5.76 15.639-17.609 15.639h-1V60h1c14.99 0 23.246-7.522 23.246-21.18V22.333l7.508 7.939l3.728-3.413L40.184 12z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M40.184 12L26.13 26.859l3.729 3.413l7.51-7.944V38.82c0 10.524-5.76 15.639-17.609 15.639h-1V60h1c14.99 0 23.246-7.522 23.246-21.18V22.333l7.508 7.939l3.728-3.413L40.184 12z"/>`),
		g.Group(children),
	)
}