package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 19c4.658 0 8.5-2.161 8.5-5S17.658 9 13 9c-4.658 0-8.5 2.161-8.5 5s3.842 5 8.5 5Zm0-9c4.179 0 7.5 1.868 7.5 4c0 2.132-3.321 4-7.5 4s-7.5-1.868-7.5-4c0-2.132 3.321-4 7.5-4Z" clip-rule="evenodd"/><path d="M12.5 6.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0v-3Zm4.01.402a.5.5 0 0 1 .98.196l-.5 2.5a.5.5 0 0 1-.98-.196l.5-2.5Zm-7.02 0a.5.5 0 0 0-.98.196l.5 2.5a.5.5 0 0 0 .98-.196l-.5-2.5ZM5.429 8.243a.5.5 0 0 0-.858.514l1.5 2.5a.5.5 0 0 0 .858-.514l-1.5-2.5Zm15.142 0a.5.5 0 1 1 .858.514l-1.5 2.5a.5.5 0 1 1-.858-.514l1.5-2.5ZM16 13.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}