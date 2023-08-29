package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JagranTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.303 20.387c1.08 2.158 1.786 4.337 2.032 6.808m-10.287.897H4.598c.226-10.968 9.188-19.79 20.21-19.79c7.915 0 14.768 4.55 18.086 11.176l-9.62 4.792"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.642 28.005c.204-9.862 8.262-17.68 18.172-17.68c4.554 0 8.716 1.674 11.905 4.44l-5.766 6.566"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.076 27.926a12.904 12.904 0 0 1-.018-.684c0-7.035 5.703-12.738 12.738-12.738c1.327 0 2.606.203 3.808.58l-1.167 3.992"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.754 27.927a9.05 9.05 0 0 1 18.078-.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 32.697a7 7 0 1 1-14 0a7 7 0 0 1 14 0Zm-9.6-4h5.2m-2.6 8v-8"/>`),
		g.Group(children),
	)
}