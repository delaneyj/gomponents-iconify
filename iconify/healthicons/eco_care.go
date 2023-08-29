package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcoCare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="m24.433 28.279l-1.867-.707c.5-1.32 1.35-2.752 2.69-4.241l1.487 1.338c-1.183 1.315-1.901 2.53-2.31 3.61Z"/><path fill-rule="evenodd" d="M15.563 7C10.035 7 6 12.64 6 18.724C6 32.304 24 41 24 41s18-9.256 18-22.276C42 12.642 37.965 7 32.437 7c-3.835 0-6.68 2.531-8.437 6.121C22.243 9.531 19.398 7 15.562 7ZM17 24.959c0 2.807 3.142 3.703 5.441 2.957c-1.098 3.251-.078 5.784.83 6.764l1.467-1.36c-.364-.392-1.2-2.014-.509-4.422c1.207 2.237 6.771 2.519 6.771-1.7V16c-2.882 2.439-5.961 3.403-8.495 4.197C19.323 21.193 17 21.92 17 24.959Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}