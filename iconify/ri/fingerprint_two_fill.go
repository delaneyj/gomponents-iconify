package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1a9 9 0 0 1 9 9v4a8.989 8.989 0 0 1-3.811 7.355c.527-1.692.811-3.49.811-5.355v-2.001h-2V16l-.003.315a15.93 15.93 0 0 1-1.431 6.315a9.055 9.055 0 0 1-3.576.314A12.925 12.925 0 0 0 13 16V9h-2v7l-.004.288a10.95 10.95 0 0 1-2.088 6.167a8.988 8.988 0 0 1-2.625-1.503A7.962 7.962 0 0 0 8 16v-6l.005-.2a3.98 3.98 0 0 1 .549-1.832L7.109 6.523A5.973 5.973 0 0 0 6 10v6l-.004.225a5.966 5.966 0 0 1-1.12 3.272A8.952 8.952 0 0 1 3 14v-4a9 9 0 0 1 9-9Zm0 3c-1.296 0-2.496.41-3.476 1.11l1.444 1.444A4 4 0 0 1 16 10v2h2v-2a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}