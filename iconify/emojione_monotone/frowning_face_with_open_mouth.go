package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrowningFaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5c15.163 0 27.5 12.336 27.5 27.5S47.163 59.5 32 59.5z"/><circle cx="20.5" cy="26" r="5" fill="currentColor"/><circle cx="43.5" cy="26" r="5" fill="currentColor"/><path fill="currentColor" d="M31.998 38c-8.568 0-12.213 2.408-13.715 6c-.801 1.919.346 3.999.346 3.999C19.078 49.217 20.844 50 32 50c11.143 0 12.922-.783 13.369-2.001c0 0 1.146-2.08.346-3.999c-1.5-3.592-5.147-6-13.717-6m9.973 4.965l-.172.563c-.08.258-.342.472-.586.472h-18.43c-.24 0-.502-.214-.58-.472l-.174-.563c-.078-.26.008-.616.191-.797c0 0 2.232-2.168 9.777-2.168c7.547 0 9.779 2.168 9.779 2.168c.187.181.275.537.195.797"/>`),
		g.Group(children),
	)
}