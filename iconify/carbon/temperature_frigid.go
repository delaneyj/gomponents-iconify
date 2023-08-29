package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureFrigid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.167 16.89L21.72 13l5.447-3.891l2.517.839l.632-1.897L28 7.279V5h-2v2.485l-5 3.572v-5.69l2.555-1.703L22.445 2L20 3.63L17.555 2l-1.11 1.664L19 5.367v5.69l-3-2.143v2.458L18.28 13L16 14.628v2.458l3-2.143v6.76L22.445 24l1.11-1.664L21 20.633v-5.69l5 3.572V21h2v-2.279l2.316-.773l-.632-1.897l-2.517.839zM12 23a3 3 0 0 1-6 0z"/><path fill="currentColor" d="M9 30a6.993 6.993 0 0 1-5-11.89V7a5 5 0 0 1 10 0v11.11A6.993 6.993 0 0 1 9 30ZM9 4a3.003 3.003 0 0 0-3 3v11.983l-.332.299a5 5 0 1 0 6.664 0L12 18.983V7a3.003 3.003 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}