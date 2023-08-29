package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarbonForIbmDotcom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.509 17.689l-6-3.55a.997.997 0 0 0-1.018.001l-6 3.55a.998.998 0 0 0-.491.86v6.9c0 .354.187.681.491.86l6 3.55a.989.989 0 0 0 1.018 0l6-3.55a.998.998 0 0 0 .491-.86v-6.9a1 1 0 0 0-.491-.861zM28 24.88l-5 2.958l-5-2.958v-5.76l5-2.958l5 2.958v5.76z"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="9" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M28 2H4c-1.103 0-2 .897-2 2v20c0 1.103.897 2 2 2h9v-2H4V10h24v4h2V4c0-1.103-.897-2-2-2zM4 8V4h24v4H4z"/>`),
		g.Group(children),
	)
}