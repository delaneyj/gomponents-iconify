package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10.5 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M21 14c-.265 0-.66.275-.993.553a4.857 4.857 0 0 0-1.088 1.276c-.214.367-.419.813-.419 1.171c0-.358-.205-.804-.42-1.171a4.857 4.857 0 0 0-1.087-1.276C16.661 14.275 16.265 14 16 14m5-4c-.265 0-.66-.275-.993-.553a4.857 4.857 0 0 1-1.088-1.276C18.705 7.804 18.5 7.358 18.5 7c0 .358-.205.804-.42 1.171c-.285.49-.659.918-1.087 1.276c-.332.278-.728.553-.993.553M8.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}