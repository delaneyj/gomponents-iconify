package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LikeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17.36 20H8V10c.625 0 1.208-.312 1.555-.832l2.554-3.832A3 3 0 0 1 14.606 4h.213a1 1 0 0 1 .987 1.164L15 10h3.56a2 2 0 0 1 1.962 2.392l-1.2 6A2 2 0 0 1 17.36 20Z" opacity=".16"/><path fill="currentColor" d="m15 10l-.986-.164A1 1 0 0 0 15 11v-1ZM4 10V9a1 1 0 0 0-1 1h1Zm16.522 2.392l.98.196l-.98-.196ZM6 21h11.36v-2H6v2ZM18.56 9H15v2h3.56V9Zm-2.574 1.164l.806-4.835L14.82 5l-.805 4.836l1.972.328ZM14.82 3h-.213v2h.213V3Zm-3.542 1.781L8.762 8.555l1.664 1.11L12.94 5.89l-1.664-1.11ZM7.93 9H4v2h3.93V9ZM3 10v8h2v-8H3Zm17.302 8.588l1.2-6l-1.961-.392l-1.2 6l1.961.392ZM8.762 8.555A1 1 0 0 1 7.93 9v2a3 3 0 0 0 2.496-1.336l-1.664-1.11Zm8.03-3.226A2 2 0 0 0 14.82 3v2l1.973.329ZM18.56 11a1 1 0 0 1 .981 1.196l1.961.392A3 3 0 0 0 18.56 9v2Zm-1.2 10a3 3 0 0 0 2.942-2.412l-1.961-.392a1 1 0 0 1-.98.804v2ZM14.606 3a4 4 0 0 0-3.329 1.781l1.665 1.11A2 2 0 0 1 14.605 5V3ZM6 19a1 1 0 0 1-1-1H3a3 3 0 0 0 3 3v-2Z"/><path stroke="currentColor" stroke-width="2" d="M8 10v10"/></g>`),
		g.Group(children),
	)
}