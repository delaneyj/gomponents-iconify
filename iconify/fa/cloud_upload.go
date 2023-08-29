package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 736q0-14-9-23L919 361q-9-9-23-9t-23 9L522 712q-10 12-10 24q0 14 9 23t23 9h224v352q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5V768h224q13 0 22.5-9.5t9.5-22.5zm640 288q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-130 70-240t188-165q-2-30-2-43q0-212 150-362T768 0q156 0 285.5 87T1242 318q71-62 166-62q106 0 181 75t75 181q0 76-41 138q130 31 213.5 135.5T1920 1024z"/>`),
		g.Group(children),
	)
}