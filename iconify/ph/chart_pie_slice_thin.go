package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieSliceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 109.5v-72a4 4 0 0 0-5.33-3.77a100 100 0 0 0-65 112.5A4 4 0 0 0 35.6 149L98 113a4 4 0 0 0 2-3.5Zm-8-2.31l-55.33 31.92A92 92 0 0 1 92 43.34Zm122.79-28.83c-.06-.13-.12-.25-.19-.37s-.15-.23-.22-.34A100.07 100.07 0 0 0 128 28a4 4 0 0 0-4 4v94.12l-80.8 47.06a4 4 0 0 0-1.43 5.49A100 100 0 0 0 228 128a99.42 99.42 0 0 0-13.21-49.64ZM132 36.09a92.05 92.05 0 0 1 73.59 42.51L132 121.46ZM128 220a92.47 92.47 0 0 1-77.22-42L130 131.91l.1-.06l79.5-46.32A92 92 0 0 1 128 220Z"/>`),
		g.Group(children),
	)
}