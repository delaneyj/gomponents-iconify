package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1536 96v1120q0 50-34 89t-86 60.5t-103.5 32t-96.5 10.5t-96.5-10.5t-103.5-32t-86-60.5t-34-89t34-89t86-60.5t103.5-32t96.5-10.5q105 0 192 39V526L640 763v709q0 50-34 89t-86 60.5t-103.5 32T320 1664t-96.5-10.5t-103.5-32t-86-60.5t-34-89t34-89t86-60.5t103.5-32T320 1280q105 0 192 39V352q0-31 19-56.5t49-35.5L1412 4q12-4 28-4q40 0 68 28t28 68z"/>`),
		g.Group(children),
	)
}