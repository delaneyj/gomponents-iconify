package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 800q0-14-9-23t-23-9h-224V416q0-13-9.5-22.5T992 384H800q-13 0-22.5 9.5T768 416v352H544q-13 0-22.5 9.5T512 800q0 14 9 23l352 352q9 9 23 9t23-9l351-351q10-12 10-24zm640 224q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-130 70-240t188-165q-2-30-2-43q0-212 150-362T768 0q156 0 285.5 87T1242 318q71-62 166-62q106 0 181 75t75 181q0 76-41 138q130 31 213.5 135.5T1920 1024z"/>`),
		g.Group(children),
	)
}