package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenRoundedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.75 2a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0V2Zm1.822 8.805a.75.75 0 1 0-.563 1.39c.814.33 1.466.981 1.796 1.796a.75.75 0 1 0 1.39-.563a4.764 4.764 0 0 0-2.623-2.623Z"/><path fill-rule="evenodd" d="M12 7.25A8.75 8.75 0 0 0 3.25 16v5.25H2a.75.75 0 0 0 0 1.5h20a.75.75 0 0 0 0-1.5h-1.25V16A8.75 8.75 0 0 0 12 7.25Zm.75 14h6.5V16a7.25 7.25 0 1 0-14.5 0v5.25h6.5v-1.628a2.251 2.251 0 1 1 1.5 0v1.628Zm-.75-3a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z" clip-rule="evenodd"/><path d="M21.53 5.47a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0Zm-18 0a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.06-1.06l-1.5-1.5Z"/></g>`),
		g.Group(children),
	)
}