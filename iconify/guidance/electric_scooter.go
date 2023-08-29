package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricScooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3 14a5 5 0 0 1 5 5v.5h.25S10 19 12 19s3.75.5 3.75.5H16V19a5 5 0 0 1 2.86-4.52M13 2.5h3.5v.825c0 3.852.808 7.65 2.36 11.155m0 0c.337.761.709 1.508 1.115 2.239m0 0a2.5 2.5 0 1 0 2.05 4.561a2.5 2.5 0 0 0-2.05-4.561ZM5.5 19a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		g.Group(children),
	)
}