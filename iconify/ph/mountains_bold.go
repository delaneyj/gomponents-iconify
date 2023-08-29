package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 84a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-40a8 8 0 1 1-8 8a8 8 0 0 1 8-8Zm90.32 149.88l-54.56-92.08a19.85 19.85 0 0 0-17.21-9.8a19.83 19.83 0 0 0-17.2 9.8l-18.7 31.55l-37.42-63.5a20 20 0 0 0-34.46 0L1.66 193.91A12 12 0 0 0 12 212h232a12 12 0 0 0 10.32-18.12ZM92 87.87L108.57 116H75.43ZM33 188l28.28-48h61.44l13.58 23.06v.07L151 188Zm145.86 0l-18.3-31l22-37.1l40.38 68.1Z"/>`),
		g.Group(children),
	)
}