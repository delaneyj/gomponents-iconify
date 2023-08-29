package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20 21.25V16c0-2.8 0-4.2-.545-5.27a5 5 0 0 0-2.185-2.185C16.2 8 14.8 8 12 8c-2.8 0-4.2 0-5.27.545a5 5 0 0 0-2.185 2.185C4 11.8 4 13.2 4 16v5.25h16Z" opacity=".5"/><path d="M14.25 10.815a.75.75 0 0 1 .75-.75c1.618 0 2.996 1.272 2.935 2.962a.75.75 0 0 1-1.5-.054c.029-.77-.593-1.408-1.435-1.408a.75.75 0 0 1-.75-.75ZM4 21.25H2a.75.75 0 0 0 0 1.5h20a.75.75 0 0 0 0-1.5H4Zm8.75-2.45a1.5 1.5 0 1 0-1.5 0v2.45h1.5V18.8ZM12 1.25a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-1.5 0V2a.75.75 0 0 1 .75-.75Zm9.53 4.28l-1.5 1.5a.75.75 0 1 1-1.06-1.06l1.5-1.5a.75.75 0 1 1 1.06 1.06Zm-18-1.06a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.06-1.06l-1.5-1.5Z"/></g>`),
		g.Group(children),
	)
}