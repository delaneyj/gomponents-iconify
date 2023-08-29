package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.987 25.607l-1.965 6.309l-1.964-6.309l-1.964 6.309l-1.964-6.309m10.304 5.777c.434.365.904.532 1.958.532h.534c.87 0 1.574-.706 1.574-1.577h0c0-.871-.705-1.577-1.574-1.577h-1.068c-.87 0-1.574-.707-1.574-1.578h0c0-.87.704-1.577 1.574-1.577h.534c1.054 0 1.523.167 1.958.532m-24.088 5.777v-3.928a2.38 2.38 0 0 0-2.381-2.38h0a2.38 2.38 0 0 0-2.381 2.38m0 3.928v-6.309m11.995 5.108a2.38 2.38 0 0 1-2.068 1.202h0a2.38 2.38 0 0 1-2.381-2.381v-1.548a2.38 2.38 0 0 1 2.38-2.38h0a2.38 2.38 0 0 1 2.382 2.38v.774h-4.762"/><rect width="4.762" height="6.309" x="28.594" y="16.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.381" ry="2.381"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.913 16.084v3.928a2.38 2.38 0 0 0 2.38 2.38h0a2.38 2.38 0 0 0 2.381-2.38v-3.929m0 3.929v2.381m2.616-3.929a2.38 2.38 0 0 1 2.38-2.38h0m-2.38 0v6.309m-10.34-1.202a2.38 2.38 0 0 1-2.069 1.202h0a2.38 2.38 0 0 1-2.38-2.381v-1.548a2.38 2.38 0 0 1 2.38-2.38h0a2.38 2.38 0 0 1 2.38 2.38v.774h-4.76"/>`),
		g.Group(children),
	)
}