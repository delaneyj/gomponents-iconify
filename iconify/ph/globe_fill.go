package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm78.37 64h-35.43a142.39 142.39 0 0 0-20.26-45a88.37 88.37 0 0 1 55.69 45ZM128 40.11c12 13 21 29.55 26.37 47.89h-52.74C107 69.66 116 53.13 128 40.11ZM96 128a145.29 145.29 0 0 1 2-24h60a145.72 145.72 0 0 1 0 48H98a145.29 145.29 0 0 1-2-24Zm5.63 40h52.74C149 186.34 140 202.87 128 215.89c-12-13.02-21-29.55-26.37-47.89Zm49.05 45a142.39 142.39 0 0 0 20.26-45h35.43a88.37 88.37 0 0 1-55.69 45Zm23.53-61a161.79 161.79 0 0 0 0-48h38.46a88.15 88.15 0 0 1 0 48Z"/>`),
		g.Group(children),
	)
}