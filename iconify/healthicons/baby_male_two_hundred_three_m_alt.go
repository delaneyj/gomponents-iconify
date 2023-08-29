package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMaleTwoHundredThreeMAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M33 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path fill-rule="evenodd" d="M12 22a6 6 0 0 0 4.89-9.476L20 9.414V11a1 1 0 1 0 2 0V7a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h1.586l-3.11 3.11A6 6 0 1 0 12 22Zm0-2a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path d="M33 42c-6.188 0-9-4.548-9-7.959V16.447a30.797 30.797 0 0 1 3.615 1.784a40.928 40.928 0 0 1 6.371 4.522c3.579 3.09 6.57 6.763 8.014 10.434v.854C42 38.021 39.187 42 33 42Zm2.293-20.761a43.134 43.134 0 0 0-1.26-1.048C36.49 20.186 38.957 19.178 42 17v11.83c-1.738-2.752-4.107-5.346-6.707-7.591Z"/></g>`),
		g.Group(children),
	)
}