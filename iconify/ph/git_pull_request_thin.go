package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 64a28 28 0 1 0-32 27.71v72.58a28 28 0 1 0 8 0V91.71A28 28 0 0 0 100 64Zm-48 0a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm40 128a20 20 0 1 1-20-20a20 20 0 0 1 20 20Zm112-27.71v-40.41a51.66 51.66 0 0 0-15.23-36.77L153.66 52H192a4 4 0 0 0 0-8h-48a4 4 0 0 0-4 4v48a4 4 0 0 0 8 0V57.66l35.11 35.11A43.71 43.71 0 0 1 196 123.88v40.41a28 28 0 1 0 8 0ZM200 212a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}