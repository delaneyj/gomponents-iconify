package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoopCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M20 5.5h-7a2 2 0 0 0-2 2V21a1.5 1.5 0 0 0 1.5 1.5H20a2 2 0 0 0 2-2v-13a2 2 0 0 0-2-2Z"/><path fill-rule="evenodd" d="M10 7.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3h-7.5A2.5 2.5 0 0 1 10 21V7.5Zm3-1a1 1 0 0 0-1 1V21a.5.5 0 0 0 .5.5H20a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-7Z" clip-rule="evenodd"/><path d="M13.153 17.026a3.193 3.193 0 1 1-5.178-3.737a3.193 3.193 0 0 1 5.178 3.737Z"/><path fill-rule="evenodd" d="M11.847 13.38a2.193 2.193 0 1 0-2.566 3.556a2.193 2.193 0 0 0 2.566-3.556Zm-4.683-.677a4.193 4.193 0 1 1 1.827 6.342L6.9 21.942a1 1 0 1 1-1.622-1.17l2.091-2.898a4.195 4.195 0 0 1-.205-5.17Z" clip-rule="evenodd"/></g><path d="M15.675 18.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M15 17.962a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 6.5A2.5 2.5 0 0 1 11.5 4h7A2.5 2.5 0 0 1 21 6.5v13a2.5 2.5 0 0 1-2.5 2.5H11a2 2 0 0 1-2-2v-1a.5.5 0 0 1 1 0v1a1 1 0 0 0 1 1h7.5a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 18.5 5h-7A1.5 1.5 0 0 0 10 6.5v3a.5.5 0 0 1-1 0v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.64 11.974a2.693 2.693 0 1 0-3.152 4.367a2.693 2.693 0 0 0 3.152-4.367Zm-4.57.022a3.693 3.693 0 1 1 1.258 5.421L4.994 20.65a.5.5 0 1 1-.81-.585l2.333-3.232a3.694 3.694 0 0 1-.447-4.836Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}