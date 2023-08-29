package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazonLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 168v32a12 12 0 0 1-24 0v-3.09C215.56 208.41 180.25 236 128 236c-64.6 0-103.3-42.18-104.92-44a12 12 0 1 1 17.84-16c.3.33 33.48 36 87.08 36c42.65 0 72.34-22.58 82.87-32H208a12 12 0 0 1 0-24h32a12 12 0 0 1 12 12Zm-96-81.92V84a32 32 0 0 0-58.83-17.45a12 12 0 0 1-20.11-13.1A56 56 0 0 1 180 84v92a12 12 0 0 1-23.85 1.81a56 56 0 1 1-.15-91.73Zm0 45.92a32 32 0 1 0-32 32a32 32 0 0 0 32-32Z"/>`),
		g.Group(children),
	)
}