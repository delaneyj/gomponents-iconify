package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLabelCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.474 6.474v.003a36.006 36.006 0 0 0-.115 1.815c-.058 1.46-.05 3.193.142 4.48a.386.386 0 0 0 .117.21l7.754 7.754a.5.5 0 0 0 .707 0l5.657-5.657a.5.5 0 0 0 0-.707l-7.754-7.754a.386.386 0 0 0-.21-.117c-1.287-.192-3.02-.2-4.48-.142a35.46 35.46 0 0 0-1.815.115h-.003ZM5.59 5.59c-.097.092-.489 4.567-.077 7.33c.044.294.19.56.4.77l7.753 7.754a1.5 1.5 0 0 0 2.121 0l5.657-5.657a1.5 1.5 0 0 0 0-2.121L13.689 5.91a1.383 1.383 0 0 0-.77-.4c-2.763-.411-7.238-.02-7.33.078Z"/><path d="M10.576 10.661a.5.5 0 1 0 .707.707a.5.5 0 0 0-.707-.707Zm-.707 1.414a1.5 1.5 0 1 1 2.121-2.121a1.5 1.5 0 0 1-2.12 2.121Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLabelCircleFilled0)"/></g>`),
		g.Group(children),
	)
}