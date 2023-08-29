package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FhirLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M24.65 26.361c-5.62 5.642-4.926 11.294-2.374 15.551C16.413 39.412 11 29.5 18.5 20.5C26.5 12.5 25 9 25 6c2.5 4.5 6.406 13.58-.35 20.361Z"/><path d="M23.496 42c-1.196-2.593-2.444-8.865 1.18-12.5c6.528-6.73 5.48-12.627 5.335-12.983L30 16.5c.003 0 .006.006.01.017c.269.424 5.39 8.73 1.621 17.983c2.175-1.256 3.1-2.971 3.369-3.5c-1.126 9.124-6.882 10.626-11.504 11ZM21.5 15.5c-3 1.833-9.782 6.5-10.182 12.5c-.4 6 3 8.5 5 10.5c-1-1.5-3-5.2-3-10s5.515-10.833 8.182-13Z"/></g>`),
		g.Group(children),
	)
}