package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaybackProgress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPlaybackProgress0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 5h40v26H4z"/><path fill="#000" stroke="#000" d="m22 14l6 4l-6 4v-8Z"/><path stroke="#fff" d="M11 40H6m11 0h25m-25 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPlaybackProgress0)"/>`),
		g.Group(children),
	)
}