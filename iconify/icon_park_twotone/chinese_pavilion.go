package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChinesePavilion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChinesePavilion0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M15 12h18s3.363 3.011 6 4c1.864.7 5 1 5 1s-1.816.649-3 1c-1.544.458-4 1-4 1H11s-2.456-.542-4-1c-1.184-.351-3-1-3-1s3.136-.3 5-1c2.637-.989 6-4 6-4Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m18 12l6-8l6 8H18Z"/><path stroke-linecap="round" d="M11 19v19m26-19v19"/><path stroke-linejoin="round" d="M6 38h36v6H6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChinesePavilion0)"/>`),
		g.Group(children),
	)
}