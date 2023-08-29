package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 4.765V20h2v-5h2v5h2V4.766A6.218 6.218 0 0 0 12 4c-1.089 0-2.11.277-3 .765Zm8 1.616v3.72A6.979 6.979 0 0 1 22 8h1v14H1V8h1c1.959 0 3.73.804 5 2.101v-3.72l-.033.03l-1.379-1.448l.724-.69c.35-.333.73-.636 1.136-.904A8.216 8.216 0 0 1 12 2a8.216 8.216 0 0 1 5.69 2.276l.724.69l-1.38 1.448L17 6.38ZM7 15a5.002 5.002 0 0 0-4-4.9V20h4v-5Zm10 5h4v-9.9a5.002 5.002 0 0 0-4 4.9v5ZM11 6.998h2.004v2.004H11V6.998Z"/>`),
		g.Group(children),
	)
}