package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoFSwipeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M8 6a4 4 0 0 0-4 4c0 1.863 1.276 3.4 3 3.844v5.719L5.719 18.28A1 1 0 0 0 4.78 18a1 1 0 0 0-.5 1.719l3 3a1 1 0 0 0 1.438 0l3-3a1.016 1.016 0 1 0-1.438-1.438L9 19.563v-5.72c1.724-.444 3-1.98 3-3.843a4 4 0 0 0-4-4zm10 0a4 4 0 0 0-4 4c0 1.863 1.276 3.4 3 3.844v5.719l-1.281-1.282A1 1 0 0 0 14.78 18a1 1 0 0 0-.5 1.719l3 3a1 1 0 0 0 1.438 0l3-3a1.016 1.016 0 1 0-1.438-1.438L19 19.563v-5.72c1.724-.444 3-1.98 3-3.843a4 4 0 0 0-4-4z"/>`),
		g.Group(children),
	)
}