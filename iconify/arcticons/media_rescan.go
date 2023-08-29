package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRescan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 23.954a1.749 1.749 0 0 0-.51-1.2L29.131 8.886a1.759 1.759 0 0 0-1.249-.51H6.498a1.999 1.999 0 0 0-1.998 1.95v27.278a1.998 1.998 0 0 0 1.998 1.999h21.124a1.82 1.82 0 0 0 .57 0h13.36a1.999 1.999 0 0 0 1.948-1.999Zm-5.256-1.429h-4.227m7.115 4.516h-7.115m7.115 4.527h-7.115m7.115 4.526h-7.115"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.706 22.652a7.043 7.043 0 0 0-1.956-3.687a7.136 7.136 0 0 0-10.07 0m-1.956 6.397a7.123 7.123 0 0 0 12.026 3.674"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.045 18.331l1.414 1.414l1.89 1.89h-5.194v-5.207l1.89 1.903zm9.939 9.938l-1.903-1.903h5.207v5.207l-1.89-1.89l-1.414-1.414z"/>`),
		g.Group(children),
	)
}