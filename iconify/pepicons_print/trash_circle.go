package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2" transform="translate(3 3)"><rect width="10" height="11.5" x="7" y="7" rx="1"/><path fill-rule="evenodd" d="M10.016 4.25A2.003 2.003 0 0 1 12 2.5c1.018 0 1.86.765 1.984 1.75H17a.75.75 0 0 1 0 1.5H7a.75.75 0 0 1 0-1.5h3.016Z" clip-rule="evenodd"/></g><path d="M11.5 17.999a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm-1-10.5h-3a1.501 1.501 0 0 1 3-.001Z"/><path d="M7.5 7.999a.5.5 0 1 1 0-1h11a.5.5 0 0 1 0 1h-11Z"/><path fill-rule="evenodd" d="M17.5 8.5h-9A.5.5 0 0 0 8 9v11a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V9a.5.5 0 0 0-.5-.5ZM9 19.5v-10h8v10H9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}