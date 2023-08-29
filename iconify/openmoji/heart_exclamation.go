package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiHeartExclamation0" d="M52 20.65a8.51 8.51 0 0 0-16-4.044a8.51 8.51 0 0 0-16 4.044a8.47 8.47 0 0 0 1.886 5.337l-.003.002L36 43.486l14.117-17.497l-.003-.002A8.47 8.47 0 0 0 52 20.65z"/></defs><g fill="#FFA7C0"><use href="#openmojiHeartExclamation0"/><circle cx="36.022" cy="55.007" r="5"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36.022" cy="55.007" r="5"/><use href="#openmojiHeartExclamation0" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}