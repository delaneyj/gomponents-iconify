package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparkasseGeorge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="9.341" height="8.081" x="12.061" y="35.527" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.552"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.24 38.22h7.162m-9.341 2.694h7.162"/><circle cx="16.725" cy="33.975" r="1.552" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="16.439" r="11.939" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.939 31.561A11.939 11.939 0 0 1 24 43.5"/>`),
		g.Group(children),
	)
}